package productusecase

import (
	"crud_api/src/domain/model"
	productrepository "crud_api/src/infrastructure/repository/product"
)

type ProductUsecase struct{
	repository productrepository.ProductRepository
}

func NewProductUsecase(repo productrepository.ProductRepository) ProductUsecase{
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

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error){
	product, err := pu.repository.GetProductById(id_product)
	if err != nil{
		return nil, err
	}
	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(id_product int) (string, error){
	deletedProduct, err := pu.repository.DeleteProductById(id_product)
	if err != nil{
		return "Atenção: Erro ao tentar excluir produto.", err
	}
	return deletedProduct, nil
}