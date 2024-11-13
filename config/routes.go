package config

import (
	"github.com/darkphotonKN/online-trade/internal/auth"
	"github.com/darkphotonKN/online-trade/internal/item"
	"github.com/darkphotonKN/online-trade/internal/member"
	"github.com/darkphotonKN/online-trade/internal/rating"
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

	// -- MEMBER --

	// --- Member Setup ---
	memberRepo := member.NewMemberRepository(DB)
	memberService := member.NewMemberService(memberRepo)
	memberHandler := member.NewMemberHandler(memberService, ratingService)

	// --- Member Routes ---
	memberRoutes := api.Group("/member")
	memberRoutes.GET("/:id", memberHandler.GetMemberByIdHandler)
	memberRoutes.POST("/signup", memberHandler.CreateMemberHandler)
	memberRoutes.POST("/signin", memberHandler.LoginMemberHandler)

	// -- ITEM --

	// --- Item Setup ---
	itemRepo := item.NewItemRepository(DB)
	itemService := item.NewItemService(itemRepo)
	itemHandler := item.NewItemHandler(itemService)

	// --- Item Routes ---
	itemRoutes := api.Group("/item")
	// Protected Routes
	itemRoutes.Use(auth.AuthMiddleware())
	itemRoutes.GET("/", itemHandler.GetItemsHandler)
	itemRoutes.POST("/", itemHandler.CreateItemHandler)
	itemRoutes.PATCH("/:id", itemHandler.UpdateItemsHandler)

	return router
}
