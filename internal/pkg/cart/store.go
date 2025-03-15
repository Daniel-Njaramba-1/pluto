package cart

import (
	"github.com/jmoiron/sqlx"
)

func (cart *Cart) FeedGetID() *int {
	return &cart.ID
}

func (cart *Cart) FeedCreateQuery() string {
	return `
		INSERT INTO carts (user_id, is_active)
		VALUES (:user_id, :is_active)
		RETURNING id
	`
}

func (cart *Cart) FeedGetByIDQuery() string {
	return `
		SELECT id, user_id, is_active, created_at, updated_at
		FROM carts
		WHERE id = $1
	`
}

func (cart *Cart) FeedGetAllQuery() string {
	return `
		SELECT id, user_id, is_active, created_at, updated_at
		FROM carts
		ORDER BY id ASC
	`
}

func (cart *Cart) FeedUpdateDetailsQuery() string {
	return `
		UPDATE carts
		SET user_id = :user_id
		WHERE id = :id
	`
}

func (cart *Cart) FeedDeactivateQuery() string {
	return `
		UPDATE carts
		SET is_active = false
		WHERE id = :id
	`
}

func (cart *Cart) FeedReactivateQuery() string {
	return `
		UPDATE carts
		SET is_active = true
		WHERE id = :id
	`
}

func (cart *Cart) FeedDeleteQuery() string {
	return `
		DELETE FROM carts
		WHERE id = :id
	`
}

func GetCartByCustomer(db *sqlx.DB, customer_id int, cart *Cart) error {
	query := `
		SELECT id, user_id, is_active, created_at, updated_at
		FROM carts
		WHERE user_id = $1
	`
	return db.Get(cart, query, customer_id)
}







