package controllers

import (
	"crud_api/src/application/usecase"
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