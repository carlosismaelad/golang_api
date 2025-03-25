package rootroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRootRoutes configura a rota inicial "/"
func RegisterRootRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "API OK"})
	})
}
