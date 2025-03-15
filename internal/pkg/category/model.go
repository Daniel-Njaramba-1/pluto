package category

import "time"

type Category struct {
	ID			int			`db:"id" json:"id"`
	SectionID   int  		`db:"section_id" json:"section_id"`
	Name		string		`db:"name" json:"name"`
	Description	string		`db:"description" json:"description"`
	IsActive	bool		`db:"is_active" json:"is_active"`
	CreatedAt	time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt	time.Time	`db:"updated_at" json:"updated_at"`
}

