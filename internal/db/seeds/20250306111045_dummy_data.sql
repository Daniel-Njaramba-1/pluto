-- +goose Up
-- +goose StatementBegin
INSERT INTO brands (name, description, is_active)
VALUES 
    ('Coca-Cola', 'Global Beverages Corporation', TRUE),
    ('PepsiCo', 'Global Food and Beverage Corporation', TRUE),
    ('Red Bull', 'Energy Drink Manufacturer', TRUE),
    ('Nestlé', 'Global Food and Beverage Company', TRUE);

INSERT INTO sections (name, description, is_active)
VALUES 
    ('Drinks', 'Beverages of various types', TRUE),
    ('Food', 'Edible products', TRUE),
    ('Snacks', 'Light meals or refreshments', TRUE);

INSERT INTO categories (section_id, name, description, is_active)
VALUES 
    ((SELECT id FROM sections WHERE name = 'Drinks' LIMIT 1), 'Soft Drinks', 'Carbonated sweet beverages', TRUE),
    ((SELECT id FROM sections WHERE name = 'Drinks' LIMIT 1), 'Energy Drinks', 'Caffeinated and stimulating beverages', TRUE),
    ((SELECT id FROM sections WHERE name = 'Drinks' LIMIT 1), 'Juices', 'Extracted liquid from fruits or vegetables', TRUE),
    ((SELECT id FROM sections WHERE name = 'Food' LIMIT 1), 'Confectionery', 'Sweet and sugary food items', TRUE),
    ((SELECT id FROM sections WHERE name = 'Snacks' LIMIT 1), 'Chips and Crisps', 'Thinly sliced fried or baked potato snacks', TRUE);

INSERT INTO products (category_id, brand_id, name, description, is_active)
VALUES 
    ((SELECT id FROM categories WHERE name = 'Soft Drinks' LIMIT 1), (SELECT id FROM brands WHERE name = 'Coca-Cola' LIMIT 1), 'Coca-Cola Classic', 'Original Coke flavor', TRUE),
    ((SELECT id FROM categories WHERE name = 'Soft Drinks' LIMIT 1), (SELECT id FROM brands WHERE name = 'Coca-Cola' LIMIT 1), 'Fanta Orange', 'Orange-flavored soft drink', TRUE),
    ((SELECT id FROM categories WHERE name = 'Soft Drinks' LIMIT 1), (SELECT id FROM brands WHERE name = 'PepsiCo' LIMIT 1), 'Pepsi Cola', 'Pepsi''s flagship cola', TRUE),
    ((SELECT id FROM categories WHERE name = 'Energy Drinks' LIMIT 1), (SELECT id FROM brands WHERE name = 'Red Bull' LIMIT 1), 'Red Bull Original', 'Stimulating energy drink', TRUE),
    ((SELECT id FROM categories WHERE name = 'Confectionery' LIMIT 1), (SELECT id FROM brands WHERE name = 'Nestlé' LIMIT 1), 'KitKat Bar', 'Crispy wafer bar covered in milk chocolate', TRUE),
    ((SELECT id FROM categories WHERE name = 'Chips and Crisps' LIMIT 1), (SELECT id FROM brands WHERE name = 'PepsiCo' LIMIT 1), 'Lay''s Potato Chips', 'Thinly sliced potato crisps', TRUE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM products;
DELETE FROM categories;
DELETE FROM sections;
DELETE FROM brands;
-- +goose StatementEnd