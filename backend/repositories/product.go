package repositories

import (
	"shopping-cart/backend/data"
	"shopping-cart/backend/models"
)

// ProductRepository handles product data access
type ProductRepository struct {
	products []models.Product
}

// NewProductRepository creates a new product repository with initial data
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: data.Products,
	}
}

// GetAll returns all products
func (r *ProductRepository) GetAll() []models.Product {
	return r.products
}

// GetByID returns a product by ID
func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}
