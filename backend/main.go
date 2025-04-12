package main

import (
	"context"
	"log"
	"shopping-cart/backend/config"
	"shopping-cart/backend/handlers"
	"shopping-cart/backend/middleware"
	"shopping-cart/backend/repositories"
	"shopping-cart/backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Shopping Cart API
// @version         1.0
// @description     A shopping cart API with product and order management
// @host            localhost:8080
// @BasePath        /api

func initializeRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password
		DB:       0,  // Default DB
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v. Discount codes will not be persistent.", err)
		return nil
	}

	return rdb
}

func initializeDependencies() (*handlers.ProductHandler, *handlers.OrderHandler, *handlers.DiscountHandler, *services.ProductService) {
	redisClient := initializeRedis()

	productRepo := repositories.NewProductRepository()

	productService := services.NewProductService(productRepo)
	discountService := services.NewDiscountService(redisClient)
	orderService := services.NewOrderService(productService, discountService)

	if redisClient != nil {
		ctx := context.Background()
		if err := discountService.InitializeDiscountCodes(ctx); err != nil {
			log.Printf("Failed to initialize discount codes: %v", err)
		} else {
			log.Println("Discount codes initialized successfully")
		}
	}

	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)
	discountHandler := handlers.NewDiscountHandler(discountService)

	return productHandler, orderHandler, discountHandler, productService
}

// setupRoutes configures all API routes
func setupRoutes(
	r *gin.Engine,
	productHandler *handlers.ProductHandler,
	orderHandler *handlers.OrderHandler,
	discountHandler *handlers.DiscountHandler,
	productService *services.ProductService,
) {
	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", middleware.InjectService(productService), handlers.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
		}

		orders := api.Group("/orders")
		{
			orders.POST("", middleware.InjectService(productService), handlers.PlaceOrder)
		}

		discounts := api.Group("/discounts")
		{
			discounts.POST("/validate", discountHandler.ValidateDiscountCode)
		}
	}
}

func main() {
	config.InitializeSwagger()

	r := gin.Default()

	r.Use(cors.New(config.ConfigureCORS()))

	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	productHandler, orderHandler, discountHandler, productService := initializeDependencies()

	setupRoutes(r, productHandler, orderHandler, discountHandler, productService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
