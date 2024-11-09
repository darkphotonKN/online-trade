package config

import (
	"github.com/darkphotonKN/online-trade/internal/rating"
	"github.com/darkphotonKN/online-trade/internal/user"
	"github.com/gin-gonic/gin"
)

/**
* Sets up API prefix route and all routers.
**/
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- RATING --

	// --- Rating Setup ---
	ratingRepo := rating.NewRatingRepository(DB)
	ratingService := rating.NewRatingService(ratingRepo)

	// -- USER --

	// --- User Setup ---
	userRepo := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService, ratingService)

	// --- User Routes ---
	userRoutes := api.Group("/user")
	userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	userRoutes.POST("/signup", userHandler.CreateUserHandler)
	userRoutes.POST("/signin", userHandler.LoginUserHandler)

	return router
}
