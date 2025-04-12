package config

import "shopping-cart/backend/docs"

func InitializeSwagger() {
	docs.SwaggerInfo.Title = "Shopping Cart API"
	docs.SwaggerInfo.Description = "A shopping cart API with product and order management"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
