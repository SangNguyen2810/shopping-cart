package models

// ErrorResponse represents a detailed error response
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e *ErrorResponse) Error() string {
	return e.Message
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

// Common error codes
const (
	// Client errors (4xx)
	ErrCodeNotFound          = "NOT_FOUND"
	ErrCodeInvalidRequest    = "INVALID_REQUEST"
	ErrCodeValidationError   = "VALIDATION_ERROR"
	ErrCodeUnauthorized      = "UNAUTHORIZED"
	ErrCodeForbidden         = "FORBIDDEN"
	ErrCodeRateLimitExceeded = "RATE_LIMIT_EXCEEDED"
	ErrCodeInvalidInput      = "INVALID_INPUT"
	ErrCodeDuplicateEntry    = "DUPLICATE_ENTRY"
	ErrCodeInvalidToken      = "INVALID_TOKEN"
	ErrCodeExpiredToken      = "EXPIRED_TOKEN"

	// Server errors (5xx)
	ErrCodeInternalError      = "INTERNAL_ERROR"
	ErrCodeDatabaseError      = "DATABASE_ERROR"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
	ErrCodeNetworkError       = "NETWORK_ERROR"

	// Business logic errors
	ErrCodeInsufficientStock = "INSUFFICIENT_STOCK"
	ErrCodeInvalidDiscount   = "INVALID_DISCOUNT"
	ErrCodeOrderNotFound     = "ORDER_NOT_FOUND"
	ErrCodePaymentFailed     = "PAYMENT_FAILED"
)
