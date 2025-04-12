package handlers

import (
	"net/http"
	"shopping-cart/backend/models"
	"shopping-cart/backend/services"

	"github.com/gin-gonic/gin"
)

// DiscountHandler handles discount code-related HTTP requests
type DiscountHandler struct {
	service *services.DiscountService
}

// NewDiscountHandler creates a new discount handler
func NewDiscountHandler(service *services.DiscountService) *DiscountHandler {
	return &DiscountHandler{service: service}
}

// ValidateDiscountRequest represents a discount code validation request
type ValidateDiscountRequest struct {
	Code string `json:"code" binding:"required"`
}

// ValidateDiscountResponse represents a discount code validation response
type ValidateDiscountResponse struct {
	Valid        bool    `json:"valid"`
	DiscountRate float64 `json:"discountRate,omitempty"`
	Error        string  `json:"error,omitempty"`
}

// @Summary      Validate a discount code
// @Description  Check if a discount code is valid and not used
// @Tags         discounts
// @Accept       json
// @Produce      json
// @Param        code  body      ValidateDiscountRequest  true  "Discount code"
// @Success      200   {object}  ValidateDiscountResponse
// @Failure      400   {object}  models.APIResponse
// @Failure      500   {object}  models.APIResponse
// @Router       /discounts/validate [post]
func (h *DiscountHandler) ValidateDiscountCode(c *gin.Context) {
	var request ValidateDiscountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error: &models.ErrorResponse{
				Code:    models.ErrCodeInvalidRequest,
				Message: "Invalid request",
				Details: err.Error(),
			},
		})
		return
	}

	if request.Code == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error: &models.ErrorResponse{
				Code:    models.ErrCodeInvalidInput,
				Message: "Discount code cannot be empty",
			},
		})
		return
	}

	// Validate the discount code
	discountRate, err := h.service.ValidateDiscountCode(c.Request.Context(), request.Code)
	if err != nil {
		c.JSON(http.StatusOK, ValidateDiscountResponse{
			Valid: false,
			Error: err.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, ValidateDiscountResponse{
		Valid:        true,
		DiscountRate: discountRate,
	})
}
