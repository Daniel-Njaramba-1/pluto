package payment

import "github.com/jmoiron/sqlx"

func (payment *Payment) FeedGetID() *int {
	return &payment.ID
}

func (payment *Payment) FeedCreateQuery() string {
	return `
		INSERT INTO payments (order_id, payment_method, amount, status, transaction_id, created_at, updated_at)
		VALUES (:order_id, :payment_method, :amount, :status, :transaction_id, :created_at, :updated_at)
		RETURNING id
	`
}

func (payment *Payment) FeedGetByIDQuery() string {
	return `
		SELECT id, order_id, payment_method, amount, status, transaction_id, created_at, updated_at
		FROM payments
		WHERE id = $1
	`
}

func (payment *Payment) FeedGetAllQuery() string {
	return `
		SELECT id, order_id, payment_method, amount, status, transaction_id, created_at, updated_at
		FROM payments
	`
}

func (payment *Payment) FeedUpdateDetailsQuery() string {
	return `
		UPDATE payments
		SET order_id = :order_id, 
			payment_method = :payment_method,
			amount = :amount,
			status = :status,
			transaction_id = :transaction_id
		WHERE id = :id
		RETURNING id, order_id, payment_method, amount, status, transaction_id, created_at, updated_at
	`
}

func (payment *Payment) FeedDeleteQuery() string {
	return `
		DELETE FROM payments
		WHERE id = $1
	`
}

func CompletePaymentQuery(db *sqlx.DB, id int, payment *Payment) error {
	query := `	
		UPDATE payments
		SET status = 'completed'
		WHERE id = $1
	`
	_, err := db.NamedExec(query, payment)
	return err
}

func FailedPaymentQuery(db *sqlx.DB, id int, payment *Payment) error {
	query := `	
		UPDATE payments
		SET status = 'failed'
		WHERE id = $1
	`
	_, err := db.NamedExec(query, payment)
	return err
}