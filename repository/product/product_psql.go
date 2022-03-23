package productRepository

import (
	"database/sql"
	"store_products/model"
)

type ProductRepository struct{}

func (p ProductRepository) GetProducts(db *sql.DB, product model.Product, ids string, products []model.Product) ([]model.Product, error) {

	rows, err := db.Query("select * from products where store_id=$1", ids)

	if err != nil {
		return []model.Product{}, err
	}

	for rows.Next() {

		if errx := rows.Scan(&product.ID, &product.Store_id, &product.Product_id, &product.Available); errx != nil {
			return []model.Product{}, err
		}

		products = append(products, product)
	}

	return products, err

}

func (p ProductRepository) AddProducts(db *sql.DB, product model.Product) (int, error) {
	err := db.QueryRow("insert into products (id,store_id, product_id, available) values($1, $2, $3,$4) RETURNING product_id;", product.ID, product.Store_id, product.Product_id, product.Available).Scan(&product.ID)

	if err != nil {
		return 0, err
	}
	return product.ID, nil
}
