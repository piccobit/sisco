package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
	"sisco/ent/tag"
	"sisco/internal/auth"
	"sisco/internal/cfg"
	"sisco/internal/db"
	"sisco/internal/exit"
	"sisco/internal/rpc/srpc"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// type gRPCServer struct

type Login struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
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

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	var err error

	dbConn, err = db.New()
	if err != nil {
		exit.Fatalf(1, "failed opening connection to database: %v", err)
	}
	defer dbConn.Close()

	grpcListenAddr := fmt.Sprintf(":%d", cfg.Config.GRPCPort)

	grpcSrv, err := srpc.New(
		srpc.ListenAddr(grpcListenAddr),
		srpc.UseTLS(cfg.Config.UseTLS),
		srpc.TLSCertFile(cfg.Config.TLSCertFile),
		srpc.TLSKeyFile(cfg.Config.TLSKeyFile),
	)

	go rpcServer(grpcSrv)

	listenAddr := fmt.Sprintf("[::]:%d", cfg.Config.Port)

	router := setupAPIRouter()

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

	go heartbeatChecker(quit)

	<-quit

	log.Println("Shutdown servers...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("HTTP server is shutting down...")

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("gRPC server is shutting down...")

	grpcSrv.GracefulStop()

	// Catching ctx.Done(). Timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("Exiting...")
	}
}

func httpServer(srv *http.Server) {
	log.Printf("HTTP server listening at %s", srv.Addr)

	if cfg.Config.UseTLS {
		if err := srv.ListenAndServeTLS(cfg.Config.TLSCertFile, cfg.Config.TLSKeyFile); err != nil && err != http.ErrServerClosed {
			exit.Fatalf(1, "failed listening: %v", err)
		}
	} else {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			exit.Fatalf(1, "failed listening: %v", err)
		}
	}

	log.Println("HTTP server shut down")
}

func rpcServer(srv *srpc.Server) {
	srv.Run()

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

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
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

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
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

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
		c.IndentedJSON(http.StatusOK, s)
	} else {
		c.JSON(http.StatusOK, s)
	}
}

func apiDeleteService(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")
	paramService := c.Param("service")

	err := dbConn.DeleteService(ctx, paramService, paramArea)
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

func apiListServices(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")

	s, err := dbConn.QueryServices(ctx, paramArea, "")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
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

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
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

func apiRegisterService(c *gin.Context) {
	var err error
	var rs RegisterService

	ctx := context.Background()

	areaParam := c.Param("area")
	serviceParam := c.Param("service")

	err = c.ShouldBindJSON(&rs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	bearer := c.Request.Header.Get("Bearer")

	token, err := dbConn.QueryToken(ctx, bearer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = dbConn.CreateService(ctx, serviceParam, areaParam, token.User, rs.Description, rs.Protocol, rs.Host, rs.Port, rs.Tags)
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

	authToken, isAdminToken, err := dbConn.QueryAuthToken(ctx, l.User, l.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
		c.IndentedJSON(http.StatusOK, gin.H{
			"token":        authToken,
			"isAdminToken": isAdminToken,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token":        authToken,
			"isAdminToken": isAdminToken,
		})
	}
}

func apiHeartbeat(c *gin.Context) {
	ctx := context.Background()

	paramService := c.Param("service")
	paramArea := c.Param("area")

	err := dbConn.UpdateServiceAvailableHeartbeat(ctx, paramService, paramArea, true, time.Now())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	prettyQuery, _ := strconv.ParseBool(c.Query("pretty"))

	if pretty || prettyQuery {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func checkPermissions(permissions auth.Permissions) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		bearer := c.Request.Header.Get("Bearer")

		if len(bearer) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			token, err := dbConn.QueryAuthTokenInfo(ctx, bearer, permissions)
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if !token.IsValid {
					c.AbortWithStatus(http.StatusUnauthorized)
				} else if (permissions & auth.ServiceOwnerOnly) != 0 {
					paramService := c.Param("service")
					paramArea := c.Param("area")

					se, err := dbConn.QueryService(ctx, paramService, paramArea)
					if err != nil {
						c.AbortWithStatus(http.StatusUnauthorized)
					}

					if !strings.EqualFold(token.Requester, se.Owner) {
						c.AbortWithStatus(http.StatusUnauthorized)
					} else {
						c.Next()
					}
				} else {
					c.Next()
				}
			}
		}
	}
}

func heartbeatChecker(done chan os.Signal) {
	t := time.NewTicker(time.Second * time.Duration(cfg.Config.HeartbeatCheckInSeconds))
	defer t.Stop()

	for {
		select {
		case <-done:
			return
		case <-t.C:
			// Check heartbeats
			log.Println("Checking heartbeats...")
			checkHeartbeats()
		}
	}
}

func checkHeartbeats() {
	ctx := context.Background()

	l, err := dbConn.QueryServices(ctx, "", "")
	if err != nil {
		exit.Fatalf(1, "could not query services: %v", err)
	}

	for _, e := range l {
		if e.Available {
			if int(time.Now().Sub(e.Heartbeat).Seconds()) > cfg.Config.HeartbeatCheckInSeconds {
				inArea, err := e.Edges.AreaOrErr()
				if err != nil {
					exit.Fatalf(1, "could not get service area: %v", err)
				}

				err = dbConn.UpdateServiceAvailable(ctx, e.Name, inArea.Name, e.Owner, false)
				if err != nil {
					exit.Fatalf(1, "could not update service availability: %w", err)
				}

				log.Println(fmt.Sprintf("Heartbeat for service '%s' is missing", e.Name))
			}
		}
	}
}

func setupAPIRouter() *gin.Engine {
	if cfg.Config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	err := router.SetTrustedProxies(cfg.Config.TrustedProxies)
	if err != nil {
		exit.Fatalf(1, "failed setting trusted proxies: %v", err)
	}

	v1Group := router.Group("/api/v1")
	v1Group.POST("/login", apiLogin)

	listGroup := v1Group.Group("/list", checkPermissions(auth.Admin|auth.Service|auth.User))
	listGroup.GET("/areas", apiListAreas)
	listGroup.GET("/service/:service/in/:area", apiGetServiceInArea)
	listGroup.GET("/services/in/:area", apiListServices)
	listGroup.GET("/services/with/:tag", apiListServicesWithTag)
	listGroup.GET("/tags", apiListTags)

	adminGroup := v1Group.Group("/admin", checkPermissions(auth.Admin|auth.Service))
	adminGroup.PUT("/heartbeat/:service/:area", apiHeartbeat)

	registerGroup := adminGroup.Group("/register")
	registerGroup.POST("/area/:area", apiRegisterArea, checkPermissions(auth.Admin))
	registerGroup.POST("/service/:service/in/:area", apiRegisterService)

	deleteGroup := adminGroup.Group("/delete")
	deleteGroup.DELETE("/service/:service/in/:area", apiDeleteService, checkPermissions(auth.ServiceOwnerOnly))
	deleteGroup.DELETE("/area/:area", apiDeleteArea)
	deleteGroup.DELETE("/tag/:tag", apiDeleteTag)

	return router
}
