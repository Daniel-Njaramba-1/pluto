package cart

import "time"

type Cart struct {
	ID				int			`db:"id" json:"id"`
	CustomerID		int			`db:"customer_id" json:"customer_id"`
	IsActive		bool		`db:"is_active" json:"is_active"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}






