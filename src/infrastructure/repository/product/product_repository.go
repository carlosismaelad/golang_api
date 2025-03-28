package productrepository

import (
	"crud_api/src/domain/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct{
	connection *sql.DB
}

func NewProductRepository(connection * sql.DB) ProductRepository{
	return ProductRepository{
		connection: connection,
	}
}

func(pr *ProductRepository) GetProducts() ([]model.Product, error){
	query := "SELECT id, product_name, price FROM product;"
	rows, err := pr.connection.Query(query)
	if err != nil{
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next(){
		err = rows.Scan(
			&productObject.ID,
			&productObject.Name,
			&productObject.Price)
		if err != nil{
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObject)
	}
	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error){
	var id int
	query, err := pr.connection.Prepare(
		"INSERT INTO product(product_name, price) VALUES ($1, $2) RETURNING id;")
	if err != nil{
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil

}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error){
	query, err := pr.connection.Prepare(
		"SELECT * FROM product where id = $1")
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	var product model.Product
	
	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}

func (pr *ProductRepository) DeleteProductById(id_product int) (string, error){
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil{
		fmt.Println(err)
		return "Atenção: Erro ao realizar consulta para exclusão.", err
	}
	defer query.Close()

	result, err := query.Exec(id_product)
	if err != nil{
		fmt.Println(err)	
		return "Atenção: produto para exclusão não localizado", err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return "Atenção: Erro ao verificar a exclusão do produto.", err
	}

	if rowsAffected == 0 {
		return "Atenção: produto para exclusão não localizado", nil
	}

	return "Produto excluído com sucesso!", nil	
}