package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products []Product
var idCounter int = 1

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	var product Product
	err := json.NewDecoder(body).Decode(&product)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	product.ID = idCounter
	products = append(products, product)
	idCounter++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(product); err != nil {
				return
			}
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var updatedProduct Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == updatedProduct.ID {
			products[i] = updatedProduct

	http.Error(w, "Product not found", http.StatusNotFound)
}
