package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/produts", nil).Methods("GET")
	r.HandleFunc("/api/v1/produts/{id}", nil).Methods("GET")
	r.HandleFunc("/api/v1/produts", nil).Methods("POST")
	r.HandleFunc("/api/v1/produts/{id}", nil).Methods("PUT")
	r.HandleFunc("/api/v1/produts/{id}", nil).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}
