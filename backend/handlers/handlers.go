package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"shopping-cart/backend/models"

	"shopping-cart/backend/services"

	"github.com/gin-gonic/gin"
)

// Create a random number generator
var rng = rand.New(rand.NewSource(1))

func GetProducts(c *gin.Context) {
	// Get products from service layer
	products := c.MustGet("productService").(*services.ProductService).GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get product service from context
	productService := c.MustGet("productService").(*services.ProductService)
	products := productService.GetAllProducts()

	// Create order items with full product details
	var orderItems []gin.H
	for _, item := range order.Items {
		for _, product := range products {
			if product.ID == item.ProductID {
				orderItems = append(orderItems, gin.H{
					"product":  product,
					"quantity": item.Quantity,
				})
				break
			}
		}
	}

	// Generate random order ID
	orderID := fmt.Sprintf("ORD%d", rng.Intn(100000))

	// Return OrderID and products information
	c.JSON(http.StatusOK, gin.H{
		"orderId": orderID,
		"items":   orderItems,
	})
}
