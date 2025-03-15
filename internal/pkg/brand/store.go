package brand

import (
	"pluto/internal/pkg/product"

	"github.com/jmoiron/sqlx"
)

func (brand *Brand) FeedGetID() *int {
	return &brand.ID
}   

func (brand *Brand) FeedCreateQuery() string {
	return `
		INSERT INTO brands (name, description, is_active)
		VALUES (:name, :description, :is_active)
		RETURNING id
	`   
}

func (brand *Brand) FeedGetByIDQuery() string {
	return `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM brands
		WHERE id = $1
	`
}

func (brand *Brand) FeedGetAllQuery() string {
	return `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM brands
		ORDER BY id ASC
	`
}

func (brand *Brand) FeedUpdateDetailsQuery() string {
	return `    
		UPDATE brands
		SET name = :name, 
            description = :description
		WHERE id = :id
	`
}

func (brand *Brand) FeedDeactivateQuery() string {
	return `
		UPDATE brands
		SET is_active = false
		WHERE id = :id
	`
}

func (brand *Brand) FeedReactivateQuery() string {
	return `
		UPDATE brands
		SET is_active = true
		WHERE id = :id
	`
}

func (brand *Brand) FeedDeleteQuery() string {
	return `
		DELETE FROM brands
		WHERE id = :id
	`
}

// SearchBrandsByName searches brands with a similar name.
func SearchBrandsByName(db *sqlx.DB, name string, brands *[]Brand) error {
    query := `
        SELECT id, name, description, is_active, created_at, updated_at
        FROM brands
        WHERE name ILIKE $1
    `
    return db.Select(brands, query, "%"+name+"%")
}

// GetAllProductsByBrand retrieves all products for a brand.
func GetAllProductsByBrand(db *sqlx.DB, brandID int, products *[]product.Product) error {
    query := `	
        SELECT id, sub_category_id, brand_id, name, description, is_active, created_at, updated_at
        FROM products 
        WHERE brand_id = $1
    `
    return db.Select(products, query, brandID)
}