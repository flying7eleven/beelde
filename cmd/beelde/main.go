package main

import (
	"beelde/internal/pkg/collections"
	"beelde/internal/pkg/health"
	"beelde/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
)

func createRouter(log *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(utils.GinLogger(log), gin.Recovery())
	router.Use(cors.AllowAll())

	v1 := router.Group("/v1")
	{
		v1.GET("/health", health.HealthStatus)
		v1.GET("/collections", collections.GetKnownCollections)
	}

	return router
}

func main() {
	// initialize our logging instance
	logger := logrus.New()

	// get the actual router for the whole API
	router := createRouter(logger)

	// start serving the API
	_ = router.Run(":8080")

}
