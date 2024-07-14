package product

import "DDD-HEX/internal/domain"

type ProductService interface {
	CreateProduct(name, description string, price float64) error
	GetProductByID(id string) (domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
}
