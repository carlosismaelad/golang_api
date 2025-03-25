package main

import (
	productcontrollers "crud_api/src/application/controllers/product"
	productusecase "crud_api/src/application/usecase/product"
	"crud_api/src/infrastructure/db"
	productrepository "crud_api/src/infrastructure/repository/product"
	registerroutes "crud_api/src/infrastructure/routes"

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
	ProductRepository := productrepository.NewProductRepository(dbConnection)

	//Usecase
	ProductUsecase := productusecase.NewProductUsecase(ProductRepository)

	//Controllers layer
	ProductController := productcontrollers.NewProductController(ProductUsecase)


	registerroutes.RegisterRoutes(server, &ProductController)
	server.Run(":8000")
}