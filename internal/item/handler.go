package item

import (
	"fmt"
	"net/http"

	"github.com/darkphotonKN/online-trade/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemHandler struct {
	Service *ItemService
}

func NewItemHandler(service *ItemService) *ItemHandler {
	return &ItemHandler{
		Service: service,
	}
}

func (h *ItemHandler) CreateItemHandler(c *gin.Context) {
	userId, _ := c.Get("userId")

	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error with parsing payload as JSON.")})
		return
	}

	err := h.Service.CreateItemService(userId.(uuid.UUID), item)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode": http.StatusInternalServerError, "message": fmt.Sprintf("Error when attempting to create item: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"statusCode": http.StatusCreated, "message": "Successfully created item."})
}

func (h *ItemHandler) GetItemsHandler(c *gin.Context) {
	userId, _ := c.Get("userId")

	items, err := h.Service.GetItemsService(userId.(uuid.UUID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to retrieve all items from user id: %s, \n error: %s\n", userId, err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Successfully retrieved all items.", "result": items})
}
