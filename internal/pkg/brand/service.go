package brand

import (
    "errors"
    "pluto/internal/lib/generics"
    "pluto/internal/pkg/product"

    "github.com/jmoiron/sqlx"
)

// BrandService provides all brand-related operations
type BrandService struct {
    db *sqlx.DB
}

// NewBrandService creates a new brand service
func NewBrandService(db *sqlx.DB) *BrandService {
    return &BrandService{db: db}
}

// Create creates a new brand
func (s *BrandService) Create(brand *Brand) (int, error) {
    if brand.Name == "" {
        return 0, errors.New("brand name cannot be empty")
    }
    
    return generics.CreateModel(s.db, brand)
}

// GetByID retrieves a brand by its ID
func (s *BrandService) GetByID(id int) (*Brand, error) {
    brand := &Brand{}
    err := generics.SelectModelByID(s.db, id, brand)
    if err != nil {
        return nil, err
    }
    
    return brand, nil
}

// GetAll retrieves all brands
func (s *BrandService) GetAll() ([]*Brand, error) {
    var brands []*Brand
    err := generics.SelectAllModels(s.db, &brands)
    if err != nil {
        return nil, err
    }
    
    return brands, nil
}

// Update updates a brand's details
func (s *BrandService) Update(brand *Brand) error {
    if brand.ID == 0 {
        return errors.New("brand ID cannot be zero")
    }
    
    if brand.Name == "" {
        return errors.New("brand name cannot be empty")
    }
    
    return generics.UpdateModelDetails(s.db, brand)
}

// Deactivate sets a brand's is_active field to false
func (s *BrandService) Deactivate(id int) error {
    brand := &Brand{ID: id}
    return generics.DeactivateModel(s.db, brand)
}

// Reactivate sets a brand's is_active field to true
func (s *BrandService) Reactivate(id int) error {
    brand := &Brand{ID: id}
    return generics.ReactivateModel(s.db, brand)
}

// Delete permanently deletes a brand
func (s *BrandService) Delete(id int) error {
    // Check if brand has any products
    var products []product.Product
    err := GetAllProductsByBrand(s.db, id, &products)
    if err != nil {
        return err
    }
    
    if len(products) > 0 {
        return errors.New("cannot delete brand with existing products")
    }
    
    brand := &Brand{ID: id}
    return generics.DeleteModel(s.db, brand)
}

// SearchByName searches for brands with similar names
func (s *BrandService) SearchByName(name string) ([]Brand, error) {
    var brands []Brand
    err := SearchBrandsByName(s.db, name, &brands)
    if err != nil {
        return nil, err
    }
    
    return brands, nil
}

// GetProducts retrieves all products for a brand
func (s *BrandService) GetProducts(brandID int) ([]product.Product, error) {
    var products []product.Product
    err := GetAllProductsByBrand(s.db, brandID, &products)
    if err != nil {
        return nil, err
    }
    
    return products, nil
}