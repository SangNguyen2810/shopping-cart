package models

type CartItem struct {
	ProductID string `json:"productId" validate:"required,uuid"`
	Quantity  int    `json:"quantity" validate:"required,min=1,max=100"`
}

type Order struct {
	Items        []CartItem `json:"items" validate:"required,min=1,max=10,dive"`
	DiscountCode string     `json:"discountCode,omitempty" validate:"omitempty,min=3,max=20,alphanum"`
}

type OrderConfirmation struct {
	OrderID   string     `json:"orderId" validate:"required,uuid"`
	Items     []CartItem `json:"items" validate:"required,min=1,dive"`
	Subtotal  float64    `json:"subtotal" validate:"required,min=0"`
	Discount  float64    `json:"discount" validate:"min=0"`
	Total     float64    `json:"total" validate:"required,min=0"`
	CreatedAt string     `json:"createdAt" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Status    string     `json:"status" validate:"required,oneof=pending confirmed processing shipped delivered cancelled"`
}
