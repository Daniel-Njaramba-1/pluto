package cart_item

import "time"

type CartItem struct {
	ID				int			`db:"id" json:"id"`
	CartID			int			`db:"cart_id" json:"cart_id"`
	ProductID		int			`db:"product_id" json:"product_id"`
	Quantity		int			`db:"quantity" json:"quantity"`
	IsActive		bool		`db:"is_active" json:"is_active"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}