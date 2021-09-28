package tests

import (
	"bytes"
	pcontrollers "kanbersky/api/controllers/products/v1"
	dbConfig "kanbersky/infrastructure"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	productDb := dbConfig.ConnectProductDB()
	productController := pcontrollers.ProductController(productDb)

	req, err := http.NewRequest("GET", "/api/v1/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productController.GetProducts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetProduct(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	productDb := dbConfig.ConnectProductDB()
	productController := pcontrollers.ProductController(productDb)

	q := req.URL.Query()
	q.Add("id", "2")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productController.GetProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateProduct(t *testing.T) {

	productDb := dbConfig.ConnectProductDB()
	productController := pcontrollers.ProductController(productDb)

	var jsonStr = []byte(`{"id":123,"name":"test","price": 12,"quantity": 1}`)

	req, err := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productController.CreateProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestUpdateProduct(t *testing.T) {

	productDb := dbConfig.ConnectProductDB()
	productController := pcontrollers.ProductController(productDb)

	var jsonStr = []byte(`{"id":123,"name":"test","price": 12,"quantity": 1}`)
	req, err := http.NewRequest("PUT", "/api/v1/products/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productController.UpdateProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteProduct(t *testing.T) {
	productDb := dbConfig.ConnectProductDB()
	productController := pcontrollers.ProductController(productDb)

	req, err := http.NewRequest("DELETE", "/api/v1/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "2")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productController.DeleteProduct)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}
