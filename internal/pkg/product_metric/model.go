package productmetric

import "time"

type ProductMetric struct {
	ID				int			`db:"id" json:"id"`
	ProductID		int			`db:"product_id" json:"product_id"`
	AverageRating	float32		`db:"average_rating" json:"average_rating"`
	ReviewCount		int			`db:"review_count" json:"review_count"`
	WishlistCount	int			`db:"wishlist_count" json:"wishlist_count"`
	BasePrice		float32		`db:"base_price" json:"base_price"`
	AdjustedPrice	float32		`db:"adjusted_price" json:"adjusted_price"`
	IsActive		bool		`db:"is_active" json:"is_active"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}