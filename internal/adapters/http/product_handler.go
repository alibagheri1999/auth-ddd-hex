package http

import (
	"DDD-HEX/internal/application/services/product"
	"DDD-HEX/internal/domain/DTO"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	ProductService product.ProductService
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req DTO.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.ProductService.CreateProduct(req.Name, req.Description, req.Price); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	product, err := h.ProductService.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
