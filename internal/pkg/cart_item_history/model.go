package cart_item_history

import "time"

type CartItemHistory struct {
	ID          int       `db:"id" json:"id"`
	CartID      int       `db:"cart_id" json:"cart_id"`
	OrderItemID int       `db:"order_item_id" json:"order_item_id"`
	ProductID   int       `db:"product_id" json:"product_id"`
	Quantity    int       `db:"quantity" json:"quantity"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}