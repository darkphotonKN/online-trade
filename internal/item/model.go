package item

type UpdateItemReq struct {
	Name          string  `db:"name" json:"name"`
	Description   string  `db:"description" json:"description"`
	PricePerUnit  float64 `db:"price_per_unit" json:"pricePerUnit"`
	StockQuantity int     `db:"stock_quantity" json:"stockQuantity"`
}
