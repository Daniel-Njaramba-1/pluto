Demand Indicators
-- Product view count (would need to add a product_views table)
SELECT product_id, COUNT(*) as view_count 
FROM product_views 
WHERE created_at > NOW() - INTERVAL '30 days' 
GROUP BY product_id;

-- Sales velocity
SELECT product_id, SUM(quantity)/30 as daily_sales_rate
FROM sales
WHERE created_at > NOW() - INTERVAL '30 days'
GROUP BY product_id;

Inventory Metrics
-- Calculate inventory ratio
SELECT p.id as product_id, s.quantity::decimal / NULLIF(s.stock_threshold, 0) as inventory_ratio
FROM products p
JOIN stocks s ON p.id = s.product_id;

Customer Engagement
-- Wishlist popularity
SELECT product_id, COUNT(*) as wishlist_count
FROM wishlist_items
GROUP BY product_id;

-- Review sentiment (you might need to add a sentiment score column to reviews)
SELECT product_id, AVG(rating) as avg_rating
FROM reviews
GROUP BY product_id;

Conversion Rates
-- Cart abandonment rate
SELECT ci.product_id, 
       COUNT(DISTINCT ci.id) as cart_adds,
       COUNT(DISTINCT cih.id) as purchases,
       COALESCE(COUNT(DISTINCT cih.id)::decimal / NULLIF(COUNT(DISTINCT ci.id), 0), 0) as conversion_rate
FROM cart_items ci
LEFT JOIN cart_items_history cih ON ci.id = cih.cart_item_id
GROUP BY ci.product_id;


Key Implementation Points:

Feature Engineering:

You're calculating multiple features that influence price
Normalizing features within categories for better comparison
Using a weighted approach to calculate demand


Model Training Strategy:

Create separate models per category for more precise pricing
Train on historical sales data to learn optimal pricing patterns
Store coefficients in the database for persistence


Price Adjustment Logic:

Apply changes only when price difference exceeds a threshold
Implement price constraints to avoid unreasonable prices
Track adjustment history with confidence scores

package pricing

import (
	"database/sql"
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
	"github.com/sajari/regression"
)

// Feature names for our model
const (
	FeatureBias             = "bias"
	FeatureDemandScore      = "demand_score"
	FeatureCompetitiveIndex = "competitive_index"
	FeatureSeasonalityFactor = "seasonality_factor"
	FeatureInventoryRatio   = "inventory_ratio"
	FeatureDaysInStock      = "days_in_stock"
	FeatureViewToPurchase   = "view_to_purchase_ratio"
)

// PricingFeature represents a product's feature data for pricing
type PricingFeature struct {
	ProductID          int
	DemandScore        float64
	CompetitiveIndex   float64
	SeasonalityFactor  float64
	InventoryRatio     float64
	DaysInStock        int
	ViewToPurchaseRatio float64
	BasePrice          float64
	OptimalPrice       float64
}

// PriceModel manages the regression model for dynamic pricing
type PriceModel struct {
	db          *sql.DB
	categoryID  int  // zero means global model
	regression  *regression.Regression
	lastTrained time.Time
	modelVersion string
}

// NewPriceModel creates a new pricing model for a category or globally
func NewPriceModel(db *sql.DB, categoryID int) *PriceModel {
	r := new(regression.Regression)
	r.SetObserved("price")
	
	// Add our features
	r.SetVar(0, FeatureBias) // Intercept term
	r.SetVar(1, FeatureDemandScore)
	r.SetVar(2, FeatureCompetitiveIndex)
	r.SetVar(3, FeatureSeasonalityFactor)
	r.SetVar(4, FeatureInventoryRatio)
	r.SetVar(5, FeatureDaysInStock)
	r.SetVar(6, FeatureViewToPurchase)
	
	return &PriceModel{
		db:          db,
		categoryID:  categoryID,
		regression:  r,
		modelVersion: fmt.Sprintf("v1-%d-%s", categoryID, time.Now().Format("20060102")),
	}
}

// LoadTrainingData fetches historical data for training
func (pm *PriceModel) LoadTrainingData() ([]PricingFeature, error) {
	// Filter by category if not a global model
	categoryFilter := ""
	args := []interface{}{}
	
	if pm.categoryID > 0 {
		categoryFilter = "AND p.category_id = $1"
		args = append(args, pm.categoryID)
	}
	
	query := fmt.Sprintf(`
		SELECT 
			p.id,
			COALESCE(pf.demand_score, 0) as demand_score,
			COALESCE(pf.competitive_index, 0) as competitive_index,
			COALESCE(pf.seasonality_factor, 1) as seasonality_factor,
			COALESCE(pf.inventory_ratio, 1) as inventory_ratio,
			COALESCE(pf.days_in_stock, 0) as days_in_stock,
			COALESCE(pf.view_to_purchase_ratio, 0) as view_to_purchase_ratio,
			pm.base_price,
			COALESCE(s.sale_price, pm.base_price) as price
		FROM products p
		JOIN product_metrics pm ON p.id = pm.product_id
		LEFT JOIN pricing_features pf ON p.id = pf.product_id
		LEFT JOIN (
			SELECT product_id, AVG(sale_price) as sale_price
			FROM sales
			WHERE created_at > NOW() - INTERVAL '90 days'
			GROUP BY product_id
		) s ON p.id = s.product_id
		WHERE p.is_active = true %s
	`, categoryFilter)
	
	rows, err := pm.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying training data: %w", err)
	}
	defer rows.Close()
	
	var features []PricingFeature
	for rows.Next() {
		var f PricingFeature
		if err := rows.Scan(
			&f.ProductID,
			&f.DemandScore,
			&f.CompetitiveIndex,
			&f.SeasonalityFactor,
			&f.InventoryRatio,
			&f.DaysInStock,
			&f.ViewToPurchaseRatio,
			&f.BasePrice,
			&f.OptimalPrice,
		); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		features = append(features, f)
	}
	
	return features, nil
}

// TrainModel builds the regression model
func (pm *PriceModel) TrainModel() error {
	features, err := pm.LoadTrainingData()
	if err != nil {
		return err
	}
	
	if len(features) < 10 {
		return fmt.Errorf("insufficient data points (%d) for reliable model training", len(features))
	}
	
	// Clear any existing data
	pm.regression.ClearObserved()
	
	// Add data points to the regression
	for _, f := range features {
		pm.regression.Train(
			regression.DataPoint(f.OptimalPrice, []float64{
				1.0, // Bias term
				f.DemandScore,
				f.CompetitiveIndex,
				f.SeasonalityFactor,
				f.InventoryRatio,
				float64(f.DaysInStock),
				f.ViewToPurchaseRatio,
			}),
		)
	}
	
	// Run the regression
	err = pm.regression.Run()
	if err != nil {
		return fmt.Errorf("regression failed: %w", err)
	}
	
	// Save coefficients to database
	err = pm.saveCoefficients()
	if err != nil {
		return err
	}
	
	pm.lastTrained = time.Now()
	return nil
}

// saveCoefficients stores the model parameters in the database
func (pm *PriceModel) saveCoefficients() error {
	tx, err := pm.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Delete existing coefficients for this category/global model
	query := `DELETE FROM price_model_coefficients WHERE product_category_id = $1`
	_, err = tx.Exec(query, pm.categoryID)
	if err != nil {
		return err
	}
	
	// Insert new coefficients
	query = `
		INSERT INTO price_model_coefficients 
		(feature_name, coefficient, product_category_id, last_trained_at) 
		VALUES ($1, $2, $3, $4)
	`
	
	for i, name := range []string{
		FeatureBias,
		FeatureDemandScore,
		FeatureCompetitiveIndex,
		FeatureSeasonalityFactor,
		FeatureInventoryRatio,
		FeatureDaysInStock,
		FeatureViewToPurchase,
	} {
		coef := pm.regression.Coeff(i)
		_, err = tx.Exec(query, name, coef, pm.categoryID, pm.lastTrained)
		if err != nil {
			return err
		}
	}
	
	return tx.Commit()
}

// LoadCoefficients retrieves stored model parameters
func (pm *PriceModel) LoadCoefficients() error {
	query := `
		SELECT feature_name, coefficient 
		FROM price_model_coefficients 
		WHERE product_category_id = $1
		ORDER BY id
	`
	
	rows, err := pm.db.Query(query, pm.categoryID)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	coefficients := make(map[string]float64)
	for rows.Next() {
		var name string
		var coef float64
		if err := rows.Scan(&name, &coef); err != nil {
			return err
		}
		coefficients[name] = coef
	}
	
	// If we have no coefficients, no model exists yet
	if len(coefficients) == 0 {
		return fmt.Errorf("no model exists for category %d", pm.categoryID)
	}
	
	// Create a new regression with these coefficients
	r := new(regression.Regression)
	r.SetObserved("price")
	
	features := []string{
		FeatureBias,
		FeatureDemandScore,
		FeatureCompetitiveIndex,
		FeatureSeasonalityFactor,
		FeatureInventoryRatio,
		FeatureDaysInStock,
		FeatureViewToPurchase,
	}
	
	for i, name := range features {
		r.SetVar(i, name)
		// We need to manually set the coefficients since we're not training
		r.Coeff(i) = coefficients[name]
	}
	
	pm.regression = r
	return nil
}

// PredictPrice calculates an optimal price for a product
func (pm *PriceModel) PredictPrice(product PricingFeature) (float64, float64, error) {
	prediction, err := pm.regression.Predict([]float64{
		1.0, // Bias term
		product.DemandScore,
		product.CompetitiveIndex,
		product.SeasonalityFactor,
		product.InventoryRatio,
		float64(product.DaysInStock),
		product.ViewToPurchaseRatio,
	})
	
	if err != nil {
		return 0, 0, err
	}
	
	// Calculate confidence based on R² value
	// Higher R² means more confidence in the model
	confidence := pm.regression.R2
	
	// Ensure the price doesn't go below some reasonable threshold
	// (e.g., don't go below cost or some minimum margin)
	if prediction < product.BasePrice * 0.8 {
		prediction = product.BasePrice * 0.8
	}
	
	return prediction, confidence, nil
}

// ApplyPriceChange updates a product's price in the database
func (pm *PriceModel) ApplyPriceChange(productID int, oldPrice, newPrice, confidence float64) error {
	tx, err := pm.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	
	// Update product_metrics table
	_, err = tx.Exec(`
		UPDATE product_metrics
		SET adjusted_price = $1, updated_at = NOW()
		WHERE product_id = $2
	`, newPrice, productID)
	if err != nil {
		return err
	}
	
	// Record the price adjustment
	_, err = tx.Exec(`
		INSERT INTO price_adjustments
		(product_id, old_price, new_price, model_version, confidence_score)
		VALUES ($1, $2, $3, $4, $5)
	`, productID, oldPrice, newPrice, pm.modelVersion, confidence)
	if err != nil {
		return err
	}
	
	return tx.Commit()
}

// UpdateAllPrices runs the pricing model on all relevant products
func (pm *PriceModel) UpdateAllPrices() error {
	// Load current product features
	features, err := pm.LoadProductFeatures()
	if err != nil {
		return err
	}
	
	// Ensure we have a trained model
	if pm.regression == nil {
		if err := pm.LoadCoefficients(); err != nil {
			return err
		}
	}
	
	for _, product := range features {
		currentPrice := 0.0
		err := pm.db.QueryRow(`
			SELECT adjusted_price 
			FROM product_metrics 
			WHERE product_id = $1
		`, product.ProductID).Scan(&currentPrice)
		
		if err != nil {
			return err
		}
		
		newPrice, confidence, err := pm.PredictPrice(product)
		if err != nil {
			return err
		}
		
		// Only update if the price change is significant (e.g., >2%)
		if math.Abs(newPrice-currentPrice)/currentPrice > 0.02 {
			if err := pm.ApplyPriceChange(product.ProductID, currentPrice, newPrice, confidence); err != nil {
				return err
			}
		}
	}
	
	return nil
}

// LoadProductFeatures gets current feature values for all products
func (pm *PriceModel) LoadProductFeatures() ([]PricingFeature, error) {
	// Similar to LoadTrainingData but for current features only
	categoryFilter := ""
	args := []interface{}{}
	
	if pm.categoryID > 0 {
		categoryFilter = "AND p.category_id = $1"
		args = append(args, pm.categoryID)
	}
	
	query := fmt.Sprintf(`
		SELECT 
			p.id,
			COALESCE(pf.demand_score, 0) as demand_score,
			COALESCE(pf.competitive_index, 0) as competitive_index,
			COALESCE(pf.seasonality_factor, 1) as seasonality_factor,
			COALESCE(pf.inventory_ratio, 1) as inventory_ratio,
			COALESCE(pf.days_in_stock, 0) as days_in_stock,
			COALESCE(pf.view_to_purchase_ratio, 0) as view_to_purchase_ratio,
			pm.base_price
		FROM products p
		JOIN product_metrics pm ON p.id = pm.product_id
		LEFT JOIN pricing_features pf ON p.id = pf.product_id
		WHERE p.is_active = true %s
	`, categoryFilter)
	
	rows, err := pm.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var features []PricingFeature
	for rows.Next() {
		var f PricingFeature
		if err := rows.Scan(
			&f.ProductID,
			&f.DemandScore,
			&f.CompetitiveIndex,
			&f.SeasonalityFactor,
			&f.InventoryRatio,
			&f.DaysInStock,
			&f.ViewToPurchaseRatio,
			&f.BasePrice,
		); err != nil {
			return nil, err
		}
		features = append(features, f)
	}
	
	return features, nil
}

// CalculateDemandScore computes a normalized demand metric
func CalculateDemandScore(db *sql.DB) error {
	_, err := db.Exec(`
		UPDATE pricing_features pf
		SET demand_score = subquery.demand_score,
		    updated_at = NOW()
		FROM (
			SELECT 
				p.id as product_id,
				-- Normalize sales velocity to a 0-1 scale within category
				GREATEST(LEAST(
					(s.quantity_sold / NULLIF(MAX(s.quantity_sold) OVER (PARTITION BY p.category_id), 0)),
					1
				), 0) * 0.7 +
				-- Add normalized wishlist popularity (30% weight)
				GREATEST(LEAST(
					(w.wishlist_count / NULLIF(MAX(w.wishlist_count) OVER (PARTITION BY p.category_id), 0)),
					1
				), 0) * 0.3
				as demand_score
			FROM products p
			LEFT JOIN (
				SELECT product_id, SUM(quantity) as quantity_sold
				FROM sales
				WHERE created_at > NOW() - INTERVAL '30 days'
				GROUP BY product_id
			) s ON p.id = s.product_id
			LEFT JOIN (
				SELECT product_id, COUNT(*) as wishlist_count
				FROM wishlist_items
				GROUP BY product_id
			) w ON p.id = w.product_id
		) subquery
		WHERE pf.product_id = subquery.product_id
	`)
	
	return err
}

// CalculateInventoryRatio updates inventory-related metrics
func CalculateInventoryRatio(db *sql.DB) error {
	_, err := db.Exec(`
		UPDATE pricing_features pf
		SET inventory_ratio = subquery.inventory_ratio,
		    days_in_stock = subquery.days_in_stock,
		    updated_at = NOW()
		FROM (
			SELECT 
				p.id as product_id,
				-- Inventory ratio (current stock / target threshold)
				COALESCE(s.quantity::decimal / NULLIF(s.stock_threshold, 0), 1) as inventory_ratio,
				-- Days since last stock update
				EXTRACT(DAY FROM NOW() - s.updated_at) as days_in_stock
			FROM products p
			JOIN stocks s ON p.id = s.product_id
		) subquery
		WHERE pf.product_id = subquery.product_id
	`)
	
	return err
}

// InitializeFeatures ensures all products have a feature record
func InitializeFeatures(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO pricing_features (product_id)
		SELECT p.id FROM products p
		LEFT JOIN pricing_features pf ON p.id = pf.product_id
		WHERE pf.id IS NULL
	`)
	
	return err
}



Feature Data Pipeline
// Initialize feature records for all products
if err := pricing.InitializeFeatures(db); err != nil {
    log.Fatal(err)
}

// Calculate and update feature values
if err := pricing.CalculateDemandScore(db); err != nil {
    log.Fatal(err)
}

if err := pricing.CalculateInventoryRatio(db); err != nil {
    log.Fatal(err)
}


Train Initial Model
// Global model
globalModel := pricing.NewPriceModel(db, 0)
if err := globalModel.TrainModel(); err != nil {
    log.Printf("Warning: couldn't train global model: %v", err)
}

// Per-category models
var categoryIDs []int
rows, _ := db.Query("SELECT id FROM categories WHERE is_active = true")
for rows.Next() {
    var id int
    rows.Scan(&id)
    categoryIDs = append(categoryIDs, id)
}

for _, catID := range categoryIDs {
    model := pricing.NewPriceModel(db, catID)
    if err := model.TrainModel(); err != nil {
        log.Printf("Warning: couldn't train model for category %d: %v", catID, err)
    }
}


Schedule Regular Price Updates
// Run daily or on a schedule that makes sense for your business
func schedulePriceUpdates() {
    for _, catID := range getAllCategoryIDs() {
        model := pricing.NewPriceModel(db, catID)
        if err := model.LoadCoefficients(); err != nil {
            // Fall back to global model if category model fails
            model = pricing.NewPriceModel(db, 0)
            if err := model.LoadCoefficients(); err != nil {
                log.Printf("Error: no model available for category %d", catID)
                continue
            }
        }
        
        if err := model.UpdateAllPrices(); err != nil {
            log.Printf("Error updating prices for category %d: %v", catID, err)
        }
    }
}

This approach gives you a solid foundation for implementing dynamic pricing. You can extend it by:

Adding more features such as time-based seasonality
Implementing A/B testing to validate price elasticity
Adding competition-based features by tracking competitor prices
Creating analytics dashboards to visualize pricing performance