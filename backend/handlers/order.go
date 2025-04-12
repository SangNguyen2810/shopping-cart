package handlers

import (
	"net/http"
	"shopping-cart/backend/models"
	"shopping-cart/backend/services"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// @Summary      Place a new order
// @Description  Create a new order with the provided items
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      models.Order  true  "Order details"
// @Success      200    {object}  models.APIResponse{data=models.OrderConfirmation}
// @Failure      400    {object}  models.APIResponse
// @Failure      500    {object}  models.APIResponse
// @Router       /orders [post]
func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	confirmation, err := h.service.PlaceOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, confirmation)
}
