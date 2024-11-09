package item

import (
	"github.com/darkphotonKN/online-trade/internal/models"
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
