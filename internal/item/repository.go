package item

import (
	"github.com/darkphotonKN/online-trade/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	DB *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{
		DB: db,
	}
}

func (r *ItemRepository) CreateItem(item models.Item) error {
	query := `
		INSERT INTO items(user_id, product_id, name, description, price_per_unit, stock_quantity)
	VALUES(:user_id, :product_id, :name, :description, :price_per_unit, :stock_quantity)
	`

	_, err := r.DB.NamedExec(query, item)

	if err != nil {
		return err
	}

	return nil
}

func (r *ItemRepository) GetItems(userId uuid.UUID) (*[]models.Item, error) {
	var items []models.Item

	query := `
	SELECT * FROM items
	WHERE items.user_id = $1
	`

	err := r.DB.Select(&items, query, userId)

	if err != nil {
		return nil, err
	}

	return &items, nil
}
