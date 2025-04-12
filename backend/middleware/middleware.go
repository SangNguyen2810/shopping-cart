package middleware

import (
	"net/http"
	"shopping-cart/backend/models"

	"shopping-cart/backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var apiError *models.ErrorResponse

			switch e := err.(type) {
			case *models.ErrorResponse:
				apiError = e
			default:
				apiError = &models.ErrorResponse{
					Code:    models.ErrCodeInternalError,
					Message: "Internal server error",
					Details: err.Error(),
				}
			}

			c.JSON(http.StatusInternalServerError, models.APIResponse{
				Success: false,
				Error:   apiError,
			})
		}
	}
}

func Logger() gin.HandlerFunc {
	return gin.Logger()
}

func ValidateRequest[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request T
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Error: &models.ErrorResponse{
					Code:    models.ErrCodeInvalidRequest,
					Message: "Invalid request body",
					Details: err.Error(),
				},
			})
			c.Abort()
			return
		}

		if err := validate.Struct(request); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse{
				Success: false,
				Error: &models.ErrorResponse{
					Code:    models.ErrCodeValidationError,
					Message: "Validation failed",
					Details: err.Error(),
				},
			})
			c.Abort()
			return
		}

		c.Set("validatedRequest", request)
		c.Next()
	}
}

func InjectService(service *services.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("productService", service)
		c.Next()
	}
}
