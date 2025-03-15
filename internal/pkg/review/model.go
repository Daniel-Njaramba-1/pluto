package review

import "time"

type Review struct {
	ID				int			`db:"id" json:"id"`
	CustomerID		int			`db:"customer_id" json:"customer_id"`
	ProductID		int			`db:"product_id" json:"product_id"`
	Rating			float32		`db:"rating" json:"rating"`
	ReviewText		string		`db:"review_text" json:"review_text"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}