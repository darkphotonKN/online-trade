-- Users Table
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- User Specific --
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,

    -- Trade Specific --
    average_rating DECIMAL(2, 1) DEFAULT 0 CHECK (average_rating >= 0 AND average_rating <= 5),
    response_time INTERVAL,  -- Optional: average time to respond, could be set in hours or days
    total_trades INT DEFAULT 0  -- Tracks the number of items traded or posted
);

-- automtically updating "updated_at" for users with trigger

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp ON users;
DROP FUNCTION IF EXISTS update_timestamp();

DROP TABLE IF EXISTS users;
-- +goose StatementEnd
