-- Items Table
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Links item to the user listing it
    product_id UUID NOT NULL,  -- Unique identifier for each product type
    name VARCHAR(255) NOT NULL,  -- Name of the product
    description TEXT,  -- Description of the product
    price_per_unit DECIMAL(10, 2) NOT NULL,  -- Price for each unit
    stock_quantity INT NOT NULL CHECK (stock_quantity >= 0),  -- Number of items available for sale
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- automtically updating "updated_at" for items with trigger

CREATE OR REPLACE FUNCTION update_item_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_item_timestamp
BEFORE UPDATE ON items
FOR EACH ROW
EXECUTE FUNCTION update_item_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_item_timestamp ON items;
DROP FUNCTION IF EXISTS update_item_timestamp();

DROP TABLE IF EXISTS items;
-- +goose StatementEnd


