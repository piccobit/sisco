package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"sisco/ent"
	"sisco/ent/area"
	"sisco/ent/service"
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
)

func serve() {
	var err error

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Config.DBUser,
		Config.DBPassword,
		Config.DBHost,
		Config.DBPort,
		Config.DBName,
		Config.DBSSLMode,
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

	if Config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	err = router.SetTrustedProxies(Config.TrustedProxies)
	if err != nil {
		log.Fatalf("failed setting trusted proxies: %v", err)
	}

	router.GET("/get/service/:area/:service", getService)
	router.GET("/list/areas", listAreas)
	router.GET("/list/services/:area", listServices)
	router.POST("/register/area/:area", registerArea)
	router.POST("/register/service/:area/:service", registerService)

	listenAddr := fmt.Sprintf(":%s", Config.Port)

	err = router.Run(listenAddr)
	if err != nil {
		log.Fatalf("failed listening: %v", err)
	}
}

func listAreas(c *gin.Context) {
	ctx := context.Background()

	a, err := dbClient.Area.Query().All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, a)
}

func getService(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")
	paramService := c.Param("service")

	s, err := dbClient.Service.Query().Where(service.And(service.Name(paramService), service.HasAreaWith(area.Name(paramArea)))).Only(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, s)
}

func listServices(c *gin.Context) {
	ctx := context.Background()

	paramArea := c.Param("area")

	a, err := dbClient.Area.Query().Where(area.Name(paramArea)).WithServices().Only(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, a)
}

type RegisterService struct {
	Description string `yaml:"description"`
	Protocol    string `yaml:"protocol"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}

func registerService(c *gin.Context) {
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

	s, err := dbClient.Service.Create().
		SetName(serviceParam).
		SetDescription(rs.Description).
		SetProtocol(rs.Protocol).
		SetHost(rs.Host).
		SetPort(rs.Port).
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

	_, err = dbClient.Area.Create().
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
