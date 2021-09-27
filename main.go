package main

import (
	pcontrollers "kanbersky/api/controllers/products/v1"
	dbConfig "kanbersky/infrastructure"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func initializeProductRouter() {
	r := mux.NewRouter()
	productDb := dbConfig.ConnectProductDB()

	productController := pcontrollers.ProductController(productDb)

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs1"}
	sh1 := middleware.Redoc(opts1, nil)
	r.Handle("/docs1", sh1)

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
