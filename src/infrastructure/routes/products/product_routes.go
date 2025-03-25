package productroutes

import (
	productcontrollers "crud_api/src/application/controllers/product"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, productController *productcontrollers.ProductController) {
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", productController.GetProducts)
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/:productId", productController.GetProductById)
		productRoutes.DELETE("/:productId", productController.DeleteProductById)
	}
}
