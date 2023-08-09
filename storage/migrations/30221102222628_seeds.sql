-- +goose Up

-- Insertando datos para la tabla customers
INSERT INTO customers (id, name, email, created_at, updated_at) VALUES ('a1b2c3d4-e5f6-7890-a1b2-c3d4e5f67890', 'Empresa ABC', 'contacto@empresaabc.com', NOW(), NOW());
INSERT INTO customers (id, name, email, created_at, updated_at) VALUES ('b2c3d4e5-f6a7-8901-b2c3-d4e5f6a78901', 'Empresa XYZ', 'contacto@empresaxyz.com', NOW(), NOW());

-- Insertando datos para la tabla discounts
INSERT INTO discounts (type, quantity, percentage, created_at, updated_at) VALUES ('land', 11, 5.0, NOW(), NOW());
INSERT INTO discounts (type, quantity, percentage, created_at, updated_at) VALUES ('maritime', 11, 3.0, NOW(), NOW());


-- Insertando datos para la tabla warehouses_and_ports
INSERT INTO warehouses_and_ports (id, type, name, created_at, updated_at) VALUES ('c3d4e5f6-a789-0123-c4d5-e6f7a8901234', 'land', 'Bodega Central', NOW(), NOW());
INSERT INTO warehouses_and_ports (id, type, name, created_at, updated_at) VALUES ('d4e5f6a7-8901-2345-e6f7-a89012345678', 'maritime', 'Puerto Principal', NOW(), NOW());

-- Insertando datos para la tabla product
INSERT INTO product (id, type, name, created_at, updated_at) VALUES ('e5f6a789-0123-4567-f8a9-0123456789ab', 'land', 'Producto Terrestre 1', NOW(), NOW());
INSERT INTO product (id, type, name, created_at, updated_at) VALUES ('f6a78901-2345-6789-a0bc-123456789abc', 'maritime', 'Producto Marítimo 1', NOW(), NOW());

-- Insertando datos para la tabla logistics
INSERT INTO logistics (id, customer_id, product_id, warehouse_port_id, type, quantity, registration_date, delivery_date, shipping_price, discount_shipping_price, vehicle_plate, fleet_number, guide_number, discount_percent, created_at, updated_at)
VALUES ('78901234-5678-9abc-d0ef-123456789def', 'a1b2c3d4-e5f6-7890-a1b2-c3d4e5f67890', 'e5f6a789-0123-4567-f8a9-0123456789ab', 'c3d4e5f6-a789-0123-c4d5-e6f7a8901234', 'land', 5, '2023-08-01 10:00:00', '2023-08-10 10:00:00', 100, 0, 'ABC123', NULL, 'L123456789', 0, NOW(), NOW());
INSERT INTO logistics (id, customer_id, product_id, warehouse_port_id, type, quantity, registration_date, delivery_date, shipping_price,discount_shipping_price,  vehicle_plate, fleet_number, guide_number, discount_percent, created_at, updated_at)
VALUES ('89012345-6789-abcd-ef01-23456789abcd', 'b2c3d4e5-f6a7-8901-b2c3-d4e5f6a78901', 'f6a78901-2345-6789-a0bc-123456789abc', 'd4e5f6a7-8901-2345-e6f7-a89012345678', 'maritime', 12, '2023-08-02 11:00:00', '2023-08-12 11:00:00', 100, 97, NULL, 'XYZ1234A', 'M123456789', 3, NOW(), NOW());

-- +goose Down

DELETE FROM logistics WHERE id IN ('78901234-5678-9abc-d0ef-123456789def', '89012345-6789-abcd-ef01-23456789abcd');
DELETE FROM product WHERE id IN ('e5f6a789-0123-4567-f8a9-0123456789ab', 'f6a78901-2345-6789-a0bc-123456789abc');
DELETE FROM warehouses_and_ports WHERE id IN ('c3d4e5f6-a789-0123-c4d5-e6f7a8901234', 'd4e5f6a7-8901-2345-e6f7-a89012345678');
DELETE FROM customers WHERE id IN ('a1b2c3d4-e5f6-7890-a1b2-c3d4e5f67890', 'b2c3d4e5-f6a7-8901-b2c3-d4e5f6a78901');
DELETE FROM discounts WHERE type = 'land' AND quantity = 11;
DELETE FROM discounts WHERE type = 'maritime' AND quantity = 11;
