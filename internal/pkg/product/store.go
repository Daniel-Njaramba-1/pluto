package product

import "github.com/jmoiron/sqlx"

func (product *Product) FeedGetID() *int {
	return &product.ID
}

func (product *Product) FeedCreateQuery() string {
	return `
		INSERT INTO products (category_id, brand_id, name, description, is_active)
		VALUES (:sub_category_id, :brand_id, :name, :description, :is_active)
		RETURNING id
	`
}

func (product *Product) FeedGetByIDQuery() string {
	return `
		SELECT id, category_id, brand_id, name, description, is_active, created_at, updated_at
		FROM products
		WHERE id = $1
	`
}

func (product *Product) FeedGetAllQuery() string {
	return `
		SELECT id, category_id, brand_id, name, description, is_active, created_at, updated_at
		FROM products
		ORDER BY id ASC
	`
}

func (product *Product) FeedUpdateDetailsQuery() string {
	return `
		UPDATE products
		SET category_id = :category_id, 
			brand_id = :brand_id,
			name = :name,
			description = :description
		WHERE id = :id
		RETURNING id, category_id, brand_id, name, description, is_active, created_at, updated_at
	`
}

func (product *Product) FeedDeactivateQuery() string {
	return `
		UPDATE products
		SET is_active = FALSE
		WHERE id = $1
		RETURNING id, sub_category_id, brand_id, name, description, is_active, created_at, updated_at
	`
}

func (product *Product) FeedReactivateQuery() string {
	return `
		UPDATE products
		SET is_active = TRUE
		WHERE id = $1
		RETURNING id, sub_category_id, brand_id, name, description, is_active, created_at, updated_at
	`
}

func (product *Product) FeedDeleteQuery() string {
	return `
		DELETE FROM products
		WHERE id = $1
	`
}

func SearchProductsByName(db *sqlx.DB, name string, products *[]Product) error {
	query := `	
		SELECT name, description, is_active, created_at, updated_at
		FROM products 
		WHERE name ILIKE $1
	`
	return db.Select(products, query, "%"+name+"%")
}


