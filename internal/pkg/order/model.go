package order

import "time"

type Order struct {
	ID				int			`db:"id" json:"id"`
	CustomerID		int			`db:"customer_id" json:"customer_id"`
	TotalPrice		float32		`db:"total_price" json:"total_price"`
	Status			string		`db:"status" json:"status"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	ID				int			`db:"id" json:"id"`
	OrderID			int			`db:"order_id" json:"order_id"`
	ProductID		int			`db:"product_id" json:"product_id"`
	Price			float32		`db:"price" json:"price"`
	Quantity		int			`db:"quantity" json:"quantity"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}