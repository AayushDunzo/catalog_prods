package main

import (
	"database/sql"
	"log"
	"net/http"
	"store_products/controller"
	"store_products/driver"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	controller := controller.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/store/{store_id}", controller.GetProducts(db)).Methods("GET")

	router.HandleFunc("/stores", controller.AddProducts(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
