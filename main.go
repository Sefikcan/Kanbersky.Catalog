package main

import (
	pcontrollers "kanbersky/api/controllers/products/v1"
	dbConfig "kanbersky/infrastructure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeProductRouter() {
	r := mux.NewRouter()
	productDb := dbConfig.ConnectProductDB()

	productController := pcontrollers.ProductController(productDb)

	r.HandleFunc("/api/v1/products", productController.GetProducts).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", productController.GetProduct).Methods("GET")
	r.HandleFunc("/api/v1/products", productController.CreateProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/{id}", productController.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", productController.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", r))
}

func main() {
	initializeProductRouter()
}
