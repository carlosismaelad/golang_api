package main

import (
	"crud_api/src/application/controllers"
	"crud_api/src/application/usecase"
	"crud_api/src/infrastructure/db"
	"crud_api/src/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	
	//Connection database layer
	dbConnection, err := db.ConnectDb()
	if err != nil{
		panic(err)
	}

	//Repository layer
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	//Controllers layer
	ProductController := controllers.NewProductController(ProductUsecase)


	server.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "API ok!",
		})
	})
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.Run(":8000")
}