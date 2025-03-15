package wishlist

import "time"

type Wishlist struct {
	ID				int			`db:"id" json:"id"`
	CustomerID		int			`db:"customer_id" json:"customer_id"`
	IsActive		bool		`db:"is_active" json:"is_active"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}

type WishlistItem struct {
	ID				int			`db:"id" json:"id"`
	WishlistID		int			`db:"wishlist_id" json:"wishlist_id"`
	ProductID		int			`db:"product_id" json:"product_id"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}