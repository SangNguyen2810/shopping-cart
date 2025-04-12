package handlers

import (
	"net/http"
	"shopping-cart/backend/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// @Summary      Get all products
// @Description  Get a list of all available products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.APIResponse{data=[]models.Product}
// @Failure      500  {object}  models.APIResponse
// @Router       /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products := h.service.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

// @Summary      Get product by ID
// @Description  Get a specific product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  models.APIResponse{data=models.Product}
// @Failure      404  {object}  models.APIResponse
// @Failure      500  {object}  models.APIResponse
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.service.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
