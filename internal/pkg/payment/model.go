package payment

import "time"

type Payment struct {
	ID				int			`db:"id" json:"id"`
	OrderID			int			`db:"order_id" json:"order_id"`
	PaymentMethod	string		`db:"payment_method" json:"payment_method"`	
	Amount			float32		`db:"amount" json:"amount"`
	Status			string		`db:"status" json:"status"`
	TransactionID	string		`db:"transaction_id" json:"transaction_id"`
	CreatedAt		time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt		time.Time	`db:"updated_at" json:"updated_at"`
}