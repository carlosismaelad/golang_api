package controllers

import (
	"crud_api/src/application/usecase"
	"crud_api/src/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct{
	productUsecase usecase.ProductUsecase
}

func NewProductController(pu usecase.ProductUsecase) productController{
	return productController{
		productUsecase: pu,
	}
}

func(p *productController) GetProducts(ctx *gin.Context){
	products, err := p.productUsecase.GetProducts()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context){
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.productUsecase.CreateProductUsecase(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)

}