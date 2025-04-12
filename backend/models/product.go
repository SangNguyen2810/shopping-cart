package models

type ProductImage struct {
	Thumbnail string `json:"thumbnail" validate:"required,url"`
	Mobile    string `json:"mobile" validate:"required,url"`
	Tablet    string `json:"tablet" validate:"required,url"`
	Desktop   string `json:"desktop" validate:"required,url"`
}

type Product struct {
	ID       string       `json:"id" validate:"required,uuid"`
	Name     string       `json:"name" validate:"required,min=2,max=100"`
	Price    float64      `json:"price" validate:"required,gt=0"`
	Image    ProductImage `json:"image" validate:"required"`
	Category string       `json:"category" validate:"required,oneof=Waffle Burger Pizza Pasta"`
}
