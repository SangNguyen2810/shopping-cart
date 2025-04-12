package services

import (
	"shopping-cart/backend/models"
	"shopping-cart/backend/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() []models.Product {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.repo.GetByID(id)
}
