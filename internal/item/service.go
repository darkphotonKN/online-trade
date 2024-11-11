package item

import (
	"fmt"

	"github.com/darkphotonKN/online-trade/internal/models"
	"github.com/google/uuid"
)

type ItemService struct {
	Repo *ItemRepository
}

func NewItemService(repo *ItemRepository) *ItemService {
	return &ItemService{
		Repo: repo,
	}
}

func (s *ItemService) CreateItemService(userId uuid.UUID, item models.Item) error {
	fmt.Println("Creating item with userID:", userId)
	// creating item for a specific user
	item.UserID = userId

	// generate a new product id
	newProdId := uuid.New()
	item.ProductID = newProdId
	return s.Repo.CreateItem(item)
}

func (s *ItemService) GetItemsService(userId uuid.UUID) (*[]models.Item, error) {
	return s.Repo.GetItems(userId)
}

func (s *ItemService) UpdateItemsService(userId uuid.UUID, id uuid.UUID, updateItemReq UpdateItemReq) (*models.Item, error) {
	return s.Repo.UpdateItemById(userId, id, updateItemReq)
}
