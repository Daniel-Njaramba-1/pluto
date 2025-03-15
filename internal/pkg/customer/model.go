package customer

import "time"

type Customer struct {
	ID					int			`db:"id" json:"id"`
	FirstName			string		`db:"firstname" json:"firstname"`
	LastName			string		`db:"lastname" json:"lastname"`
	Username			string		`db:"username" json:"username"`
	Email				string		`db:"email" json:"email"`
	PasswordHash		string		`db:"password_hash" json:"password_hash"`
	Token				string		`db:"token" json:"token"`
	Phone				string		`db:"phone" json:"phone"`
	Address				string		`db:"address" json:"address"`
	IsActive			bool		`db:"is_active" json:"is_active"`
	CreatedAt			time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt			time.Time	`db:"updated_at" json:"updated_at"`
}