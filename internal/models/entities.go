package models

import (
	"github.com/google/uuid"
	"time"
)

/**
* Shared entities that are imported by more than one package.
**/
type User struct {
	BaseDBDateModel
	Email         string  `db:"email" json:"email"`
	Name          string  `db:"name" json:"name"`
	Password      string  `db:"password" json:"password,omitempty"`
	AverageRating float64 `db:"average_rating"`
	ResponseTime  int     `db:"response_time"`
	TotalTrades   int     `db:"total_trades"`
}

type Rating struct {
	BaseIDModel
	UserID uuid.UUID `db:"user_id" json:"userId"`
	Rating int       `db:"rating" json:"rating"`
}

type Item struct {
	BaseDBDateModel
	UserID        uuid.UUID `db:"user_id" json:"userId"`
	ProductID     uuid.UUID `db:"product_id" json:"productId"`
	Name          string    `db:"name" json:"name"`
	Description   string    `db:"description" json:"description"`
	PricePerUnit  float64   `db:"price_per_unit" json:"pricePerUnit"`
	StockQuantity int       `db:"stock_quantity" json:"stockQuantity"`
}

/**
* Base models for default table columns.
**/

type BaseIDModel struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

type BaseDBUserModel struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UpdatedUser uuid.UUID `db:"updated_user" json:"updatedUser"`
	CreatedUser uuid.UUID `db:"created_user" json:"createdUser"`
}

type BaseDBDateModel struct {
	ID        uuid.UUID `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type BaseDBUserDateModel struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UpdatedUser uuid.UUID `db:"updated_user" json:"updatedUser"`
	CreatedUser uuid.UUID `db:"created_user" json:"createdUser"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
