-- +goose Up
-- +goose StatementBegin
CREATE TABLE logistics (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  customer_id UUID NOT NULL REFERENCES customers(id),
  product_id UUID NOT NULL REFERENCES product(id),
  warehouse_port_id UUID NOT NULL REFERENCES warehouses_and_ports(id),
  type VARCHAR(50) NOT NULL CHECK (type IN ('land', 'maritime')),
  quantity INT NOT NULL,
  registration_date TIMESTAMP NOT NULL,
  delivery_date TIMESTAMP NOT NULL,
  shipping_price FLOAT NOT NULL,
  discount_shipping_price FLOAT NOT NULL,
  vehicle_plate VARCHAR(10), -- Nullable porque es solo para Terrestre
  fleet_number VARCHAR(10),  -- Nullable porque es solo para Marítima
  guide_number VARCHAR(10) UNIQUE NOT NULL,
  discount_percent FLOAT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE logistics;
-- +goose StatementEnd
