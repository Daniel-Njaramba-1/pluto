package cart_item

import "github.com/jmoiron/sqlx"

func (cart_item *CartItem) FeedCreateQuery() string {
	return `
		INSERT INTO cart_items (cart_id, product_id, quantity, is_active)
		VALUES (:cart_id, :product_id, :quantity, :is_active)
		RETURNING id
	`
}

func (cart_item *CartItem) FeedGetID() *int {
	return &cart_item.ID
}

func (cart_item *CartItem) FeedGetAllQuery() string {
	return `
		SELECT id, cart_id, product_id, quantity, is_active, created_at, updated_at
		FROM cart_items
		ORDER BY id ASC
	`
}

func (cart_item *CartItem) FeedUpdateDetailsQuery() string {
	return `
		UPDATE cart_items
		SET quantity = :quantity
		WHERE id = :id
	`
}

func (cart_item *CartItem) FeedDeactivateQuery() string {
	return `
		UPDATE cart_items
		SET is_active = false
		WHERE id = :id
	`
}

func (cart_item *CartItem) FeedReactivateQuery() string {
	return `
		UPDATE cart_items
		SET is_active = true
		WHERE id = :id
	`
}

func (cart_item *CartItem) FeedDeleteQuery() string {
	return `
		DELETE FROM cart_items
		WHERE id = :id
	`
}

func GetCartItemsByCart(db *sqlx.DB, cart_id int, cart_items *[]CartItem) error {
	query := `
		SELECT id, cart_id, product_id, quantity, is_active, created_at, updated_at
		FROM cart_items
		WHERE cart_id = $1
		ORDER BY id ASC
	`
	return db.Select(cart_items, query, cart_id)
}