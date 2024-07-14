package repository

import "DDD-HEX/internal/domain"

type ProductRepository interface {
	Save(product domain.Product) error
	FindByID(id string) (domain.Product, error)
	FindAll() ([]domain.Product, error)
}
