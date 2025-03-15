package section

import (
	"errors"
	"pluto/internal/lib/generics"
	"pluto/internal/pkg/category"

	"github.com/jmoiron/sqlx"
)

// Service provides all section-related operations
type Service struct {
	db *sqlx.DB
}

// NewService creates a new section service
func NewService(db *sqlx.DB) *Service {
	return &Service{db: db}
}

// Create creates a new section
func (s *Service) Create(section *Section) (int, error) {
	if section.Name == "" {
		return 0, errors.New("section name cannot be empty")
	}
	
	return generics.CreateModel(s.db, section)
}

// GetByID retrieves a section by its ID
func (s *Service) GetByID(id int) (*Section, error) {
	section := &Section{}
	err := generics.SelectModelByID(s.db, id, section)
	if err != nil {
		return nil, err
	}
	
	return section, nil
}

// GetAll retrieves all sections
func (s *Service) GetAll() ([]*Section, error) {
	var sections []*Section
	err := generics.SelectAllModels(s.db, &sections)
	if err != nil {
		return nil, err
	}
	
	return sections, nil
}

// Update updates a section's details
func (s *Service) Update(section *Section) error {
	if section.ID == 0 {
		return errors.New("section ID cannot be zero")
	}
	
	if section.Name == "" {
		return errors.New("section name cannot be empty")
	}
	
	return generics.UpdateModelDetails(s.db, section)
}

// Deactivate sets a section's is_active field to false
func (s *Service) Deactivate(id int) error {
	section := &Section{ID: id}
	return generics.DeactivateModel(s.db, section)
}

// Reactivate sets a section's is_active field to true
func (s *Service) Reactivate(id int) error {
	section := &Section{ID: id}
	return generics.ReactivateModel(s.db, section)
}

// Delete permanently deletes a section
func (s *Service) Delete(id int) error {
	// Check if section has any categories
	var categories []category.Category
	err := GetAllCategoriesBySection(s.db, id, &categories)
	if err != nil {
		return err
	}
	
	if len(categories) > 0 {
		return errors.New("cannot delete section with existing categories")
	}
	
	section := &Section{ID: id}
	return generics.DeleteModel(s.db, section)
}

// SearchByName searches for sections with similar names
func (s *Service) SearchByName(name string) ([]Section, error) {
	var sections []Section
	err := SearchSectionsByName(s.db, name, &sections)
	if err != nil {
		return nil, err
	}
	
	return sections, nil
}

func (s *Service) GetCategories(sectionID int) ([]category.Category, error) {
	var categories []category.Category
	err := GetAllCategoriesBySection(s.db, sectionID, &categories)
	if err != nil {
		return nil, err 
	}
	
	return categories, nil
}
