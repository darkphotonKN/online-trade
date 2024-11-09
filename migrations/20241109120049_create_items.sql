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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd


