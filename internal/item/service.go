package item

import (
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
	// creating item for a specific user
	item.UserID = userId

	// generate a new product id
	newProdId := uuid.New()
	item.ProductID = newProdId
	return s.Repo.CreateItem(item)
}
