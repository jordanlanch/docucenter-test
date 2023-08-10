-- +goose Up
-- +goose StatementBegin
CREATE TABLE discounts (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  type VARCHAR(50) NOT NULL CHECK (type IN ('land', 'maritime')),
  quantity_from INT NOT NULL,
  quantity_to INT NOT NULL,
  percentage FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE discounts;
-- +goose StatementEnd
