package registerroutes

import (
	productcontrollers "crud_api/src/application/controllers/product"
	productroutes "crud_api/src/infrastructure/routes/products"
	rootroute "crud_api/src/infrastructure/routes/root"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, productController *productcontrollers.ProductController){
	rootroute.RegisterRootRoutes(router)
	productroutes.RegisterProductRoutes(router, productController)
}