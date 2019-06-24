package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthStatus(context *gin.Context) {
	context.Status(http.StatusNoContent)
}
