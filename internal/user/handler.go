package user

import (
	"fmt"
	"net/http"

	"github.com/darkphotonKN/online-trade/internal/models"
	"github.com/darkphotonKN/online-trade/internal/rating"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	Service       *UserService
	RatingService *rating.RatingService
}

func NewUserHandler(service *UserService, ratingService *rating.RatingService) *UserHandler {
	return &UserHandler{
		Service:       service,
		RatingService: ratingService,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error with parsing payload as JSON.")})
		return
	}

	err := h.Service.CreateUserService(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": http.StatusInternalServerError, "message": fmt.Sprintf("Error when attempting to create user: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"statusCode": http.StatusCreated, "message": "Successfully created user."})
}

func (h *UserHandler) LoginUserHandler(c *gin.Context) {
	var loginReq UserLoginRequest

	err := c.ShouldBindJSON(&loginReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when unmarshalling json payload: %s\n", err)})
		return
	}

	user, err := h.Service.LoginUserService(loginReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to login user: %s\n", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Successfully logged in.",
		// TODO : Add jwt tokens.
		// de-reference to return the user struct, not pointer
		"result": user})
}

func (h *UserHandler) GetUserByIdHandler(c *gin.Context) {
	// get id from param
	idParam := c.Param("id")

	// check that its a valid uuid
	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error with id %d, not a valid uuid.", id)})
		// return to stop flow of function after error response
		return
	}

	user, err := h.Service.GetUserByIdService(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to get user with id %d %s", id, err.Error())})

		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Successfully retreived user.",
		"result": user})
}

func (r *UserRepository) RefreshTokenHandler(c *gin.Context) {
	var refreshTokenReq RefreshTokenRequest

	err := c.ShouldBindJSON(&refreshTokenReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error with parsing payload as JSON.")})
		return
	}

}
