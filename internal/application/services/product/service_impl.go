package product

import (
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/repository"
)

type productServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: productRepo}
}

func (s *productServiceImpl) CreateProduct(name, description string, price float64) error {
	product := domain.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	return s.productRepository.Save(product)
}

func (s *productServiceImpl) GetProductByID(id string) (domain.Product, error) {
	return s.productRepository.FindByID(id)
}

func (s *productServiceImpl) GetAllProducts() ([]domain.Product, error) {
	return s.productRepository.FindAll()
}
