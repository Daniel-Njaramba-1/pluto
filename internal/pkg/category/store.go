package category

import (
    "pluto/internal/pkg/product"

    "github.com/jmoiron/sqlx"
)

func (category *Category) FeedGetID() *int {
    return &category.ID
}

func (category *Category) FeedCreateQuery() string {
    return `
        INSERT INTO categories (section_id, name, description, is_active)
        VALUES (:section_id, :name, :description, :is_active)
        RETURNING id
    `
}

func (category *Category) FeedGetByIDQuery() string {
    return `
        SELECT id, section_id, name, description, is_active, created_at, updated_at
        FROM categories
        WHERE id = $1
    `
}

func (category *Category) FeedGetAllQuery() string {
    return `
        SELECT id, section_id, name, description, is_active, created_at, updated_at
        FROM categories
        ORDER BY id ASC
    `
}

func (category *Category) FeedUpdateDetailsQuery() string {
    return `
        UPDATE categories
        SET section_id = :section_id,
            name = :name,
            description = :description,
            is_active = :is_active
        WHERE id = :id
        RETURNING id, section_id, name, description, is_active, created_at, updated_at
    `
}

func (category *Category) FeedDeactivateQuery() string {
    return `
        UPDATE categories
        SET is_active = FALSE
        WHERE id = $1
        RETURNING id, section_id, name, description, is_active, created_at, updated_at
    `
}

func (category *Category) FeedReactivateQuery() string {
    return `
        UPDATE categories
        SET is_active = TRUE
        WHERE id = $1
        RETURNING id, section_id, name, description, is_active, created_at, updated_at
    `
}

func (category *Category) FeedDeleteQuery() string {
    return `
        DELETE FROM categories
        WHERE id = $1
    `
}

func SearchCategoriesByName(db *sqlx.DB, name string, categories *[]Category) error {
    query := `	
        SELECT id, section_id, name, description, is_active, created_at, updated_at
        FROM categories 
        WHERE name ILIKE $1
    `
    return db.Select(categories, query, "%"+name+"%")
}

func GetAllProductsByCategory(db *sqlx.DB, category_id int, products *[]product.Product) error {
    query := `	
        SELECT id, category_id, brand_id, name, description, is_active, created_at, updated_at
        FROM products 
        WHERE category_id = $1
    `
    return db.Select(products, query, category_id)
}