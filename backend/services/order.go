package services

import (
	"context"
	"fmt"
	"math/rand"
	"shopping-cart/backend/models"
	"strings"
	"time"
)

// OrderService handles order-related business logic
type OrderService struct {
	productService  *ProductService
	discountService *DiscountService
}

// NewOrderService creates a new order service
func NewOrderService(productService *ProductService, discountService *DiscountService) *OrderService {
	return &OrderService{
		productService:  productService,
		discountService: discountService,
	}
}

func (s *OrderService) calculateDiscount(ctx context.Context, order models.Order, subtotal float64) (float64, error) {
	if order.DiscountCode == "" {
		return 0, nil
	}

	code := strings.ToUpper(order.DiscountCode)

	// Handle HAPPYHOURS codes with Redis
	if strings.HasPrefix(code, "HAPPYHOURS") {
		discountPercentage, err := s.discountService.ValidateDiscountCode(ctx, code)
		if err != nil {
			return 0, err
		}
		return subtotal * discountPercentage, nil
	}

	// Legacy code for BUYGETONE
	if code == "BUYGETONE" {
		lowestPrice, err := s.findLowestPrice(order.Items)
		if err != nil {
			return 0, err
		}
		return lowestPrice, nil
	}

	return 0, nil
}

func (s *OrderService) findLowestPrice(items []models.CartItem) (float64, error) {
	var lowestPrice float64 = -1
	for _, item := range items {
		product, err := s.productService.GetProductByID(item.ProductID)
		if err != nil {
			return 0, err
		}
		if product == nil {
			return 0, fmt.Errorf("product not found: %s", item.ProductID)
		}
		if lowestPrice == -1 || product.Price < lowestPrice {
			lowestPrice = product.Price
		}
	}
	return lowestPrice, nil
}

// PlaceOrder processes an order and returns the confirmation
func (s *OrderService) PlaceOrder(order models.Order) (models.OrderConfirmation, error) {
	ctx := context.Background()

	subtotal, err := s.calculateSubtotal(order.Items)
	if err != nil {
		return models.OrderConfirmation{}, err
	}

	discount, err := s.calculateDiscount(ctx, order, subtotal)
	if err != nil {
		return models.OrderConfirmation{}, err
	}

	// Generate random order ID
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderID := fmt.Sprintf("ORD%d", rng.Intn(100000))

	return models.OrderConfirmation{
		OrderID:   orderID,
		Items:     order.Items,
		Subtotal:  subtotal,
		Discount:  discount,
		Total:     subtotal - discount,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		Status:    "pending",
	}, nil
}

func (s *OrderService) calculateSubtotal(items []models.CartItem) (float64, error) {
	var subtotal float64
	for _, item := range items {
		product, err := s.productService.GetProductByID(item.ProductID)
		if err != nil {
			return 0, err
		}
		if product == nil {
			return 0, fmt.Errorf("product not found: %s", item.ProductID)
		}
		subtotal += product.Price * float64(item.Quantity)
	}
	return subtotal, nil
}
