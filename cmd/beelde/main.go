package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
)

func HealthStatus(context *gin.Context) {
	context.Status(http.StatusNoContent)
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("/v1")
	{
		v1.GET("/health", HealthStatus)
	}

	return router
}

func main() {
	// get the actual router for the whole API
	router := createRouter()

	// start serving the API
	_ = router.Run(":8080")

}
