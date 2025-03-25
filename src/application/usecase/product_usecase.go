package usecase

import (
	"crud_api/src/domain/model"
	"crud_api/src/infrastructure/repository"
)

type ProductUsecase struct{
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase{
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error){
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProductUsecase(product model.Product) (model.Product, error){
	productId, err := pu.repository.CreateProduct(product)
	if err != nil{
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}