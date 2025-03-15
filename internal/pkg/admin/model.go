package admin

import "time"

type Admin struct {
    ID				int			`db:"id" json:"id"`
    Username		string		`db:"username" json:"username"`
    Email			string		`db:"email" json:"email"`
    PasswordHash	string		`db:"password_hash" json:"-"`
    Password        string      `db:"-" json:"password"` // only for input, not saved in db
    IsActive		bool		`db:"is_active" json:"is_active"`
    CreatedAt		time.Time	`db:"created_at" json:"created_at"`
    UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
} 