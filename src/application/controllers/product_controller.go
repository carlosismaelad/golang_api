package controllers

import (
	"crud_api/src/application/usecase"
	"crud_api/src/domain/model"
	"net/http"
	"strconv"

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

func(p *productController) GetProductById(ctx *gin.Context){
	id := ctx.Param("productId")

	if id == ""{
		response := model.Response{
			Message: "Atenção: Id do produto não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil{
		response := model.Response{
			Message: "Atenção: Id inválido. O id do produto precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil{
		response := model.Response{
			Message: "Atenção: Produto não foi encontrado na base de dados.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func(p *productController) DeleteProductById(ctx *gin.Context){
	id := ctx.Param("productId")

	if id == ""{
		response := model.Response{
			Message: "Atenção: Id do produto não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil{
		response := model.Response{
			Message: "Atenção: Id inválido. O id do produto precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedProduct, err := p.productUsecase.DeleteProductById(productId)
	if err != nil {
		response := model.Response{
			Message: "Erro interno ao excluir produto",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Message: deletedProduct,
	}

	ctx.JSON(http.StatusOK, response)
}