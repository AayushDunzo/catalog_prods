package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"store_products/model"
	productRepository "store_products/repository/product"
	"store_products/utils"

	"github.com/gorilla/mux"
)

type Controller struct{}

var products []model.Product

func (c Controller) GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product model.Product
		var error model.Error
		params := mux.Vars(r)
		products = []model.Product{}
		productRepo := productRepository.ProductRepository{}
		id := (params["store_id"])
		prod, err := productRepo.GetProducts(db, product, id, products)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, prod)

	}
}

func (c Controller) AddProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product model.Product
		var productID int
		var error model.Error

		json.NewDecoder(r.Body).Decode(&product)

		if product.Store_id == "" || product.Product_id == "" || product.Available == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		productRepo := productRepository.ProductRepository{}
		productID, err := productRepo.AddProducts(db, product)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, productID)

	}
}
