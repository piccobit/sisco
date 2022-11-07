package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/tag"
	"sisco/internal/cfg"
	"sisco/internal/db"
	"sisco/internal/grpc/server"
)

// type gRPCServer struct

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
	dbConn   *db.Client
)

func serve() {
	var err error

	dbConn, err = db.New()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer dbConn.Close()

	if cfg.Config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	err = router.SetTrustedProxies(cfg.Config.TrustedProxies)
	if err != nil {
		log.Fatalf("failed setting trusted proxies: %v", err)
	}

	v1Group := router.Group("/api/v1")
	v1Group.POST("/login", apiLogin)

	getGroup := v1Group.Group("/get", checkToken(false))
	getGroup.GET("/service/:service/in/:area", apiGetServiceInArea)

	listGroup := v1Group.Group("/list", checkToken(false))
	listGroup.GET("/areas", apiListAreas)
	listGroup.GET("/services/in/:area", apiListServicesInArea)
	listGroup.GET("/services/with/:tag", apiListServicesWithTag)
	listGroup.GET("/tags", apiListTags)

	adminGroup := v1Group.Group("/admin", checkToken(true))

	registerGroup := adminGroup.Group("/register")
	registerGroup.POST("/area/:area", apiRegisterArea)
	registerGroup.POST("/service/:service/in/:area", apiRegisterServiceInArea)

	deleteGroup := adminGroup.Group("/delete")
	deleteGroup.DELETE("/service/:service/in/:area", apiDeleteServiceInArea)
	deleteGroup.DELETE("/area/:area", apiDeleteArea)
	deleteGroup.DELETE("/tag/:tag", apiDeleteTag)

	grpcListenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcSrv := grpc.NewServer()

	go grpcServer(grpcSrv, grpcListenAddr)

	listenAddr := fmt.Sprintf(":%d", cfg.Config.Port)

	httpSrv := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}

	go httpServer(httpSrv)

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall, but SIGKILL can't be caught, so we don't need adding it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown servers ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("HTTP server is shutting down ...")

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("gRPC server is shutting down ...")

	grpcSrv.GracefulStop()

	// Catching ctx.Done(). Timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("Exiting ...")
	}
}

func httpServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed listening: %v", err)
	}

	log.Println("HTTP server shut down")
}

func grpcServer(srv *grpc.Server, listenAddr string) {
	server.Run(srv, listenAddr)

	log.Println("gRPC server shut down")
}

func apiListAreas(c *gin.Context) {
	ctx := context.Background()

	a, err := dbConn.QueryAreas(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	if pretty {
		c.IndentedJSON(http.StatusOK, a)
	} else {
		c.JSON(http.StatusOK, a)
	}
}

func apiListTags(c *gin.Context) {
	ctx := context.Background()

	t, err := dbClient.Tag.Query().Order(ent.Asc(tag.FieldID)).All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	if pretty {
		c.IndentedJSON(http.StatusOK, t)
	} else {
		c.JSON(http.StatusOK, t)
	}
}

func apiGetServiceInArea(c *gin.Context) {
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

	if pretty {
		c.IndentedJSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

func apiDeleteServiceInArea(c *gin.Context) {
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

func apiDeleteArea(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")

	err := dbConn.DeleteArea(ctx, paramArea)
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{
			"status": "NOT deleted",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}

func apiDeleteTag(c *gin.Context) {
	ctx := context.Background()

	paramTag := c.Param("tag")

	_, err := dbClient.Tag.Delete().
		Where(tag.Name(paramTag)).Exec(ctx)
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

func apiListServicesInArea(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")

	s, err := dbClient.Service.Query().
		Where(service.HasAreaWith(area.Name(paramArea))).
		WithTags().
		Order(ent.Asc(service.FieldID)).
		All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	if pretty {
		c.IndentedJSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

func apiListServicesWithTag(c *gin.Context) {
	ctx := context.Background()

	paramTag := c.Param("tag")

	s, err := dbClient.Service.Query().
		Where(service.HasTagsWith(tag.Name(paramTag))).
		Order(ent.Asc(service.FieldID)).
		WithTags().
		All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	if pretty {
		c.IndentedJSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

type RegisterService struct {
	Description string   `yaml:"description"`
	Protocol    string   `yaml:"protocol"`
	Host        string   `yaml:"host"`
	Port        string   `yaml:"port"`
	Tags        []string `yaml:"tags,omitempty"`
}

func apiRegisterServiceInArea(c *gin.Context) {
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

	if pretty {
		c.IndentedJSON(http.StatusOK, gin.H{
			"area":    areaParam,
			"service": s,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"area":    areaParam,
			"service": s,
		})
	}
}

type RegisterArea struct {
	Description string `yaml:"description"`
}

func apiRegisterArea(c *gin.Context) {
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

	err = dbConn.CreateArea(ctx, areaParam, ra.Description)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

type Login struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func apiLogin(c *gin.Context) {
	var l Login

	ctx := context.Background()

	err := c.ShouldBindJSON(&l)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	authToken, isAdminToken, err := dbConn.GetSecretToken(ctx, l.User, l.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":        authToken,
		"isAdminToken": isAdminToken,
	})
}

func checkToken(isAdminToken bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		bearer := c.Request.Header.Get("Bearer")

		if len(bearer) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			tokenIsValid, err := dbConn.CheckToken(ctx, bearer, isAdminToken)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if tokenIsValid {
					c.Next()
				} else {
					c.AbortWithStatus(http.StatusUnauthorized)
				}
			}
		}
	}
}
