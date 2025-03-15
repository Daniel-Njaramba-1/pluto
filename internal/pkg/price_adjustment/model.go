package price_adjustment

import "time"

type PriceAdjustment struct {
	ID				int			`db:"id" json:"id"`
	ProductID		int			`db:"product_id" json:"product_id"`
	OldPrice		float32		`db:"old_price" json:"old_price"`
	NewPrice		float32		`db:"new_price" json:"new_price"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}