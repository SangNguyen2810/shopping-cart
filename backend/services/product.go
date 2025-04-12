package services

import (
	"shopping-cart/backend/models"
	"shopping-cart/backend/repositories"
)

// ProductService handles product-related business logic
type ProductService struct {
	repo *repositories.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetAllProducts returns all available products
func (s *ProductService) GetAllProducts() []models.Product {
	return s.repo.GetAll()
}

// GetProductByID returns a product by ID
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}
