package sale

func (sale *Sale) FeedGetID() *int {
	return &sale.ID
}

func (sale *Sale) FeedCreateQuery() string {
	return `
		INSERT INTO sales (order_item_id, product_id, sale_price, quantity)
		VALUES (:order_item_id, :product_id, :sale_price, :quantity)
		RETURNING id
	`
}

func (sale *Sale) FeedGetByIDQuery() string {
	return `
		SELECT id, order_item_id, product_id, sale_price, quantity, created_at, updated_at
		FROM sales
		WHERE id = $1
	`
}

func (sale *Sale) FeedGetAllQuery() string {
	return `
		SELECT id, order_item_id, product_id, sale_price, quantity, created_at, updated_at
		FROM sales
	`
}

func (sale *Sale) FeedUpdateDetailsQuery() string {
	return `
		UPDATE sales
		SET order_item_id = :order_item_id, 
			product_id = :product_id,
			sale_price = :sale_price,
			quantity = :quantity
		WHERE id = :id
		RETURNING id, order_item_id, product_id, sale_price, quantity, created_at, updated_at
	`
}

func (sale *Sale) FeedDeleteQuery() string {
	return `
		DELETE FROM sales
		WHERE id = $1
	`
}