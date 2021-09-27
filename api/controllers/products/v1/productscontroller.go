package controller

import (
	"encoding/json"
	"kanbersky/common/helpers"
	"kanbersky/infrastructure/entities"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var product entities.Product
var products []entities.Product

type ProductRepository struct {
	db *gorm.DB
}

func ProductController(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetProducts(w http.ResponseWriter, r *http.Request) {
	helpers.SetHeader(w)
	if result := p.db.Find(&products); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(helpers.PrepareErrorResponse("Product not found!", http.StatusNotFound))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (p *ProductRepository) GetProduct(w http.ResponseWriter, r *http.Request) {
	helpers.SetHeader(w)
	params := mux.Vars(r)

	if result := p.db.First(&product, params["id"]); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(helpers.PrepareErrorResponse("Product not found!", http.StatusNotFound))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductRepository) CreateProduct(w http.ResponseWriter, r *http.Request) {
	helpers.SetHeader(w)
	json.NewDecoder(r.Body).Decode(&product)

	if result := p.db.Create(&product); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.PrepareErrorResponse(result.Error.Error(), http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductRepository) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	helpers.SetHeader(w)

	params := mux.Vars(r)

	if result := p.db.First(&product, params["id"]); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(helpers.PrepareErrorResponse("Product not found!", http.StatusNotFound))
		return
	}

	json.NewDecoder(r.Body).Decode(&product)
	if response := p.db.Save(&product); response.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.PrepareErrorResponse(response.Error.Error(), http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (p *ProductRepository) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	helpers.SetHeader(w)

	params := mux.Vars(r)

	if result := p.db.Delete(&product, params["id"]); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.PrepareErrorResponse(result.Error.Error(), http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
