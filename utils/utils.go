package utils

import (
	"encoding/json"
	"net/http"
	"store_products/model"
)

func SendError(w http.ResponseWriter, status int, err model.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
