package config

import (
	"github.com/darkphotonKN/ecommerce-server-go/internal/product"
	"github.com/darkphotonKN/ecommerce-server-go/internal/rating"
	"github.com/darkphotonKN/ecommerce-server-go/internal/user"
	"github.com/gin-gonic/gin"
)

/**
* Sets up API prefix route and all routers.
**/
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- USER --

	// --- User Setup ---
	userRepo := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// --- User Routes ---
	userRoutes := api.Group("/user")
	userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	userRoutes.POST("/signup", userHandler.CreateUserHandler)
	userRoutes.POST("/signin", userHandler.LoginUserHandler)

	// -- RATING --

	// --- Rating Setup ---
	ratingRepo := rating.NewRatingRepository(DB)
	ratingService := rating.NewRatingService(ratingRepo)

	// -- PRODUCT --

	// --- Product Setup ---
	productRepo := product.NewProductRepository(DB)
	productService := product.NewProductService(productRepo, ratingService)
	productHandler := product.NewProductHandler(productService)

	// --- Product Routes ---
	productRoutes := api.Group("/product")
	productRoutes.GET("/", productHandler.GetProductsHandler)
	productRoutes.GET("/trending", productHandler.GetTrendingProductsHandler)
	productRoutes.GET("/:id", productHandler.GetProductByIdHandler)
	productRoutes.POST("/", productHandler.CreateProductsHandler)
	productRoutes.POST("/:id/rate", productHandler.CreateProductRating)
	return router
}
