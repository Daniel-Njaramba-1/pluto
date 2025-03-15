package product

import "time"

type Product struct {
	ID				int			`db:"id" json:"id"`
	CategoryID	    int			`db:"category_id" json:"category_id"`
	BrandID			int			`db:"brand_id" json:"brand_id"`
	Name			string		`db:"name" json:"name"`
	Description		string		`db:"description" json:"description"`
	Is_active		bool		`db:"is_active" json:"is_active"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}

