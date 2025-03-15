package stock

func (stock *Stock) FeedGetID() *int {
	return &stock.ID
}

func (stock *Stock) FeedCreateQuery() string {
	return `
		INSERT INTO stocks (product_id, quantity, stock_threshold)
		VALUES (:product_id, :quantity, :stock_threshold)
		RETURNING id
	`
}

func (stock *Stock) FeedGetByIDQuery() string {
	return `
		SELECT id, product_id, quantity, stock_threshold, created_at, updated_at
		FROM stocks
		WHERE id = $1
	`
}

func (stock *Stock) FeedGetAllQuery() string {
	return `
		SELECT id, product_id, quantity, stock_threshold, created_at, updated_at
		FROM stocks
	`
}

func (stock *Stock) FeedUpdateDetailsQuery() string {
	return `
		UPDATE stocks
		SET product_id = :product_id, 
			quantity = :quantity,
			stock_threshold = :stock_threshold
		WHERE id = :id
		RETURNING id, product_id, quantity, stock_threshold, created_at, updated_at
	`
}

func (stock *Stock) FeedDeleteQuery() string {
	return `
		DELETE FROM stocks
		WHERE id = $1
	`
}