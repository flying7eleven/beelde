package collections

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetKnownCollections(context *gin.Context) {
	resp := CollectionsResponse{
		Id:   "id",
		Name: "Fooo",
	}
	context.JSON(http.StatusOK, resp)
}
