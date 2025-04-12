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

// initializeRedis creates and returns a Redis client
func initializeRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// Test connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Redis connection failed: %v. Discount codes will not be persistent.", err)
		// Return nil to handle gracefully if Redis isn't available
		return nil
	}

	return rdb
}

// initializeDependencies creates and returns all required service and handler instances
func initializeDependencies() (*handlers.ProductHandler, *handlers.OrderHandler, *handlers.DiscountHandler, *services.ProductService) {
	// Initialize Redis
	redisClient := initializeRedis()

	// Initialize repositories
	productRepo := repositories.NewProductRepository()

	// Initialize services
	productService := services.NewProductService(productRepo)
	discountService := services.NewDiscountService(redisClient)
	orderService := services.NewOrderService(productService, discountService)

	// Initialize discount codes if Redis is available
	if redisClient != nil {
		ctx := context.Background()
		if err := discountService.InitializeDiscountCodes(ctx); err != nil {
			log.Printf("Failed to initialize discount codes: %v", err)
		} else {
			log.Println("Discount codes initialized successfully")
		}
	}

	// Initialize handlers
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
		// Product routes
		products := api.Group("/products")
		{
			products.GET("", middleware.InjectService(productService), handlers.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
		}

		// Order routes
		orders := api.Group("/orders")
		{
			orders.POST("", middleware.InjectService(productService), handlers.PlaceOrder)
		}

		// Discount routes
		discounts := api.Group("/discounts")
		{
			discounts.POST("/validate", discountHandler.ValidateDiscountCode)
		}
	}
}

func main() {
	// Initialize Swagger documentation
	config.InitializeSwagger()

	// Set release mode in production
	// gin.SetMode(gin.ReleaseMode)

	// Create router
	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(config.ConfigureCORS()))

	// Add middleware
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	// Initialize dependencies
	productHandler, orderHandler, discountHandler, productService := initializeDependencies()

	// Setup routes
	setupRoutes(r, productHandler, orderHandler, discountHandler, productService)

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	r.Run(":8080")
}
