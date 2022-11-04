package cmd

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap"
	"github.com/spf13/cobra"
	"sisco/cfg"
	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/tag"
	"sisco/ent/token"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Long:  `Start the server listening on the specified port.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

var (
	dbClient *ent.Client
	ldapConn *ldap.Conn
)

func serve() {
	var err error

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Config.DBUser,
		cfg.Config.DBPassword,
		cfg.Config.DBHost,
		cfg.Config.DBPort,
		cfg.Config.DBName,
		cfg.Config.DBSSLMode,
	)

	dbClient, err = ent.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer func(dbClient *ent.Client) {
		err := dbClient.Close()
		if err != nil {
			log.Fatalf("failed closing postgres DB connection: %v", err)
		}
	}(dbClient)

	if cfg.Config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	err = router.SetTrustedProxies(cfg.Config.TrustedProxies)
	if err != nil {
		log.Fatalf("failed setting trusted proxies: %v", err)
	}

	router.POST("/login", login)
	router.GET("/get/service/:service/in/:area", checkToken(false), getServiceInArea)
	router.GET("/list/areas", checkToken(false), listAreas)
	router.GET("/list/services/in/:area", checkToken(false), listServicesInArea)
	router.GET("/list/services/with/:tag", checkToken(false), listServicesWithTag)
	router.GET("/list/tags", checkToken(false), listTags)
	router.POST("/register/area/:area", checkToken(true), registerArea)
	router.POST("/register/service/:service/in/:area", checkToken(true), registerServiceInArea)
	router.DELETE("/delete/service/:service/in/:area", checkToken(true), deleteServiceInArea)
	router.DELETE("/delete/area/:area", checkToken(true), deleteArea)

	listenAddr := fmt.Sprintf(":%d", cfg.Config.Port)

	err = router.Run(listenAddr)
	if err != nil {
		log.Fatalf("failed listening: %v", err)
	}
}

func listAreas(c *gin.Context) {
	ctx := context.Background()

	a, err := dbClient.Area.Query().WithServices().All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, a)
}

func listTags(c *gin.Context) {
	ctx := context.Background()

	t, err := dbClient.Tag.Query().WithService().All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, t)
}

func getServiceInArea(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")
	paramService := c.Param("service")

	s, err := dbClient.Service.Query().
		Where(service.And(service.Name(paramService), service.HasAreaWith(area.Name(paramArea)))).
		WithTags().
		Only(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, s)
}

func deleteServiceInArea(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")
	paramService := c.Param("service")

	_, err := dbClient.Service.Delete().
		Where(service.And(service.Name(paramService), service.HasAreaWith(area.Name(paramArea)))).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

func deleteArea(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")

	numServices, err := dbClient.Service.Query().Where(service.HasAreaWith(area.Name(paramArea))).Count(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	if numServices > 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("area '%s' is not empty", paramArea),
		})

		return
	}

	_, err = dbClient.Area.Delete().
		Where(area.Name(paramArea)).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

func listServicesInArea(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")

	s, err := dbClient.Service.Query().Where(service.HasAreaWith(area.Name(paramArea))).WithTags().All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, s)
}

func listServicesWithTag(c *gin.Context) {
	ctx := context.Background()

	paramTag := c.Param("tag")

	s, err := dbClient.Service.Query().Where(service.HasTagsWith(tag.Name(paramTag))).WithTags().All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, s)
}

type RegisterService struct {
	Description string   `yaml:"description"`
	Protocol    string   `yaml:"protocol"`
	Host        string   `yaml:"host"`
	Port        string   `yaml:"port"`
	Tags        []string `yaml:"tags,omitempty"`
}

func registerServiceInArea(c *gin.Context) {
	var err error
	var rs RegisterService

	ctx := context.Background()

	areaParam := c.Param("area")

	var numAreas int

	if numAreas, err = dbClient.Area.Query().Where(area.Name(areaParam)).Count(ctx); numAreas == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("area '%s' does not exist, please register it first", areaParam),
		})

		return
	}

	serviceParam := c.Param("service")

	err = c.ShouldBindJSON(&rs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	var tags []*ent.Tag

	for _, tagName := range rs.Tags {
		t, _ := dbClient.Tag.Query().Where(tag.Name(tagName)).Only(ctx)
		if t == nil {
			t, err = dbClient.Tag.Create().SetName(tagName).Save(ctx)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		tags = append(tags, t)
	}

	s, err := dbClient.Service.Create().
		SetName(serviceParam).
		SetDescription(rs.Description).
		SetProtocol(rs.Protocol).
		SetHost(rs.Host).
		SetPort(rs.Port).
		AddTags(tags...).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	_, err = dbClient.Area.Update().Where(area.Name(areaParam)).AddServices(s).Save(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"area":    areaParam,
		"service": s,
	})
}

type RegisterArea struct {
	Description string `yaml:"description"`
}

func registerArea(c *gin.Context) {
	var ra RegisterArea

	ctx := context.Background()

	areaParam := c.Param("area")

	err := c.ShouldBindJSON(&ra)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	a, err := dbClient.Area.Create().
		SetName(areaParam).
		SetDescription(ra.Description).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"area": a,
	})
}

type Login struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func login(c *gin.Context) {
	var l Login

	ctx := context.Background()

	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	authToken := generateSecureToken(32)

	t, err := dbClient.Token.Query().Where(token.User(l.User)).Only(ctx)
	if t == nil {
		if ldapConn == nil {
			ldapConn, err = ldap.DialURL(cfg.Config.LdapURL)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})

				return
			}
			defer ldapConn.Close()
		}

		err = ldapConn.Bind(cfg.Config.LdapBindDN, cfg.Config.LdapBindPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		// We check first if this is an 'admin' token.

		isAdminToken := false

		filter := replace(cfg.Config.LdapFilterAdminsDN, "{user}", ldap.EscapeFilter(l.User))

		searchReq := ldap.NewSearchRequest(cfg.Config.LdapBaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"dn"}, []ldap.Control{})

		result, err := ldapConn.Search(searchReq)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		if len(result.Entries) != 0 {
			isAdminToken = true
		} else {
			filter = replace(cfg.Config.LdapFilterUsersDN, "{user}", ldap.EscapeFilter(l.User))

			searchReq = ldap.NewSearchRequest(cfg.Config.LdapBaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, filter, []string{"dn"}, []ldap.Control{})

			result, err = ldapConn.Search(searchReq)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})

				return
			}

			if len(result.Entries) == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "user not found",
				})

				return
			}
		}

		err = ldapConn.Bind(result.Entries[0].DN, l.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			return
		}

		_, err = dbClient.Token.Create().
			SetUser(l.User).
			SetToken(authToken).
			SetAdmin(isAdminToken).
			Save(ctx)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}
	} else {
		_, err = dbClient.Token.Update().Where(token.User(l.User)).SetCreated(time.Now()).Save(ctx)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		authToken = t.Token
	}

	c.JSON(http.StatusOK, gin.H{
		"token": authToken,
	})
}

func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func checkToken(isAdminToken bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		bearer := c.Request.Header.Get("Bearer")

		if len(bearer) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			t, err := dbClient.Token.Query().Where(token.Token(bearer)).Only(ctx)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if int(time.Now().Sub(t.Created).Seconds()) > cfg.Config.TokenValidInSeconds {
					c.AbortWithStatus(http.StatusUnauthorized)
				} else {
					if isAdminToken {
						if t.Admin {
							c.Next()
						} else {
							c.AbortWithStatus(http.StatusUnauthorized)
						}
					} else {
						c.Next()
					}
				}
			}
		}
	}
}

func replace(haystack string, needle string, replacement string) string {
	res := strings.Replace(
		haystack,
		needle,
		replacement,
		-1,
	)

	return res
}
