package model

type Product struct {
	ID         int    `json:id`
	Store_id   string `json:store_id`
	Product_id string `json:product_id`
	Available  string `json:available`
}
