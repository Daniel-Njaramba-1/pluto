package product

import (
	"errors"
	"pluto/internal/lib/generics"

	"github.com/jmoiron/sqlx"
)

// Service provides all product-related operations
type Service struct {
	db *sqlx.DB
}

// NewService creates a new product service
func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

// Create creates a new product
func (s *Service) Create(product *Product) (int, error) {
	if product.Name == "" {
		return 0, errors.New("product name cannot be empty")
	}

	return generics.CreateModel(s.db, product)
}

// GetByID retrieves a product by its ID
func (s *Service) GetByID(id int) (*Product, error) {
	product := &Product{}
	err := generics.SelectModelByID(s.db, id, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetAll retrieves all products
func (s *Service) GetAll() ([]*Product, error) {
	var products []*Product
	err := generics.SelectAllModels(s.db, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Update updates a product's details
func (s *Service) Update(product *Product) error {
	if product.ID == 0 {
		return errors.New("product ID cannot be zero")
	}

	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}

	return generics.UpdateModelDetails(s.db, product)
}

// Deactivate sets a product's is_active field to false
func (s *Service) Deactivate(id int) error {
	product := &Product{ID: id}
	return generics.DeactivateModel(s.db, product)
}

// Reactivate sets a product's is_active field to true
func (s *Service) Reactivate(id int) error {
	product := &Product{ID: id}
	return generics.ReactivateModel(s.db, product)
}

// Delete permanently deletes a product
func (s *Service) Delete(id int) error {
	product := &Product{ID: id}
	return generics.DeleteModel(s.db, product)
}

// SearchByName searches for products with similar names
func (s *Service) SearchByName(name string) ([]Product, error) {
	var products []Product
	err := SearchProductsByName(s.db, name, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}