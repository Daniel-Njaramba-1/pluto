package category

import (
	"errors"
	"pluto/internal/lib/generics"
	"pluto/internal/pkg/product"

	"github.com/jmoiron/sqlx"
)

// Service provides all category-related operations
type Service struct {
	db *sqlx.DB
}

// NewService creates a new category service
func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

// Create creates a new category
func (s *Service) Create(category *Category) (int, error) {
	if category.Name == "" {
		return 0, errors.New("category name cannot be empty")
	}
	
	return generics.CreateModel(s.db, category)
}

// GetByID retrieves a category by its ID
func (s *Service) GetByID(id int) (*Category, error) {
	category := &Category{}
	err := generics.SelectModelByID(s.db, id, category)
	if err != nil {
		return nil, err
	}
	
	return category, nil
}

// GetAll retrieves all categories
func (s *Service) GetAll() ([]*Category, error) {
	var categories []*Category
	err := generics.SelectAllModels(s.db, &categories)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}

// Update updates a category's details
func (s *Service) Update(category *Category) error {
	if category.ID == 0 {
		return errors.New("category ID cannot be zero")
	}
	
	if category.Name == "" {
		return errors.New("category name cannot be empty")
	}
	
	return generics.UpdateModelDetails(s.db, category)
}

// Deactivate sets a category's is_active field to false
func (s *Service) Deactivate(id int) error {
	category := &Category{ID: id}
	return generics.DeactivateModel(s.db, category)
}

// Reactivate sets a category's is_active field to true
func (s *Service) Reactivate(id int) error {
	category := &Category{ID: id}
	return generics.ReactivateModel(s.db, category)
}

// Delete permanently deletes a category
func (s *Service) Delete(id int) error {
	// Check if category has any products
	var products []product.Product
	err := GetAllProductsByCategory(s.db, id, &products)
	if err != nil {
		return err
	}
	
	if len(products) > 0 {
		return errors.New("cannot delete category with existing products")
	}
	
	category := &Category{ID: id}
	return generics.DeleteModel(s.db, category)
}

// SearchByName searches for categories with similar names
func (s *Service) SearchByName(name string) ([]Category, error) {
	var categories []Category
	err := SearchCategoriesByName(s.db, name, &categories)
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}

// GetProducts retrieves all products for a category
func (s *Service) GetProducts(categoryID int) ([]product.Product, error) {
	var products []product.Product
	err := GetAllProductsByCategory(s.db, categoryID, &products)
	if err != nil {
		return nil, err
	}
	
	return products, nil
}