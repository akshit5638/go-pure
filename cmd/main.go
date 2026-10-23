package main

import (
	"encoding/json"
	"net/http"
	"pure-go/internal/handlers"
)

type healthResponse struct {
	Status string `json:"status"`
}

func healthhandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := healthResponse{Status: "ok"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func main() {
	http.HandleFunc("GET /health", healthhandler)
	http.HandleFunc("POST /products", handlers.CreateProductHandler)
	http.HandleFunc("GET /products", handlers.GetProductsHandler)
	http.HandleFunc("GET /product/{id}", handlers.GetProductByIDHandler)
	http.HandleFunc("PUT /products/{id}", handlers.UpdateProductHandler)
	// http.HandleFunc("DELETE /products", handlers.DeleteProductHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
