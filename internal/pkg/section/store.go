// filepath: c:\Users\ADMIN\Desktop\pricewave\backend\pkg\section\store.go
package section

import (
	"pluto/internal/pkg/category"

	"github.com/jmoiron/sqlx"
)

func (section *Section) FeedGetID() *int {
	return &section.ID
}

func (section *Section) FeedCreateQuery() string {
	return `
		INSERT INTO sections (name, description, is_active)
		VALUES (:name, :description, :is_active)
		RETURNING id
	`
}

func (section *Section) FeedGetByIDQuery() string {
	return `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM sections
		WHERE id = $1
	`
}

func (section *Section) FeedGetAllQuery() string {
	return `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM sections
		ORDER BY id ASC
	`
}

func (section *Section) FeedUpdateDetailsQuery() string {
	return `
		UPDATE sections
		SET name = :name, 
			description = :description
		WHERE id = :id
	`
}

func (section *Section) FeedDeactivateQuery() string {
	return `
		UPDATE sections
		SET is_active = false
		WHERE id = :id
	`
}

func (section *Section) FeedReactivateQuery() string {
	return `
		UPDATE sections
		SET is_active = true
		WHERE id = :id
	`
}

func (section *Section) FeedDeleteQuery() string {
	return `
		DELETE FROM sections
		WHERE id = :id
	`
}

// SearchSectionsByName searches sections with a similar name.
func SearchSectionsByName(db *sqlx.DB, name string, sections *[]Section) error {
	query := `
		SELECT id, name, description, is_active, created_at, updated_at
		FROM sections
		WHERE name ILIKE $1
	`
	return db.Select(sections, query, "%"+name+"%")
}

// GetAllCategoriesBySection retrieves all categories for a section.
func GetAllCategoriesBySection(db *sqlx.DB, sectionID int, categories *[]category.Category) error {
	query := `
		SELECT id, section_id, name, description, is_active, created_at, updated_at
		FROM categories
		WHERE section_id = $1
	`
	return db.Select(categories, query, sectionID)
}