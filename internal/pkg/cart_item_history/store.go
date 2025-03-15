package cart_item_history

func (cart_item_history *CartItemHistory) FeedGetID() *int {
	return &cart_item_history.ID
}

func (cart_item_history *CartItemHistory) FeedCreateQuery() string {
	return `
		INSERT INTO cart_item_histories (cart_id, order_item_id, product_id, quantity)
		VALUES (:cart_id, :order_item_id, :product_id, :quantity)
		RETURNING id
	`
}

func (cart_item_history *CartItemHistory) FeedGetAllQuery() string {
	return `
		SELECT id, cart_id, order_item_id, product_id, quantity, created_at, updated_at
		FROM cart_item_histories
		ORDER BY id ASC
	`
}

func (cart_item_history *CartItemHistory) FeedUpdateDetailsQuery() string {
	return `
		UPDATE cart_item_histories
		SET quantity = :quantity	
		WHERE id = :id
	`
}

func (cart_item_history *CartItemHistory) FeedDeleteQuery() string {
	return `
		DELETE FROM cart_item_histories
		WHERE id = :id
	`
}