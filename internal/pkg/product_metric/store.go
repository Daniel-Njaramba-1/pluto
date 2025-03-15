package productmetric

func (productMetric *ProductMetric) FeedGetID() *int {		
	return &productMetric.ID
}

func (productMetric *ProductMetric) FeedCreateQuery() string {
	return `
		INSERT INTO product_metrics (product_id, average_rating, review_count, wishlist_count, base_price, adjusted_price, is_active)
		VALUES (:product_id, :average_rating, :review_count, :wishlist_count, :base_price, :adjusted_price, :is_active)
		RETURNING id
	`
}

func (productMetric *ProductMetric) FeedGetByIDQuery() string {
	return `
		SELECT id, product_id, average_rating, review_count, wishlist_count, base_price, adjusted_price, is_active, created_at, updated_at
		FROM product_metrics
		WHERE id = $1
	`
}

func (productMetric *ProductMetric) FeedGetAllQuery() string {
	return `
		SELECT id, product_id, average_rating, review_count, wishlist_count, base_price, adjusted_price, is_active, created_at, updated_at
		FROM product_metrics
	`
}

func (productMetric *ProductMetric) FeedUpdateDetailsQuery() string {
	return `
		UPDATE product_metrics
		SET product_id = :product_id, 
			average_rating = :average_rating,
			review_count = :review_count,
			wishlist_count = :wishlist_count,
			base_price = :base_price,
			adjusted_price = :adjusted_price,
			is_active = :is_active
		WHERE id = :id
		RETURNING id, product_id, average_rating, review_count, wishlist_count, base_price, adjusted_price, is_active, created_at, updated_at
	`
}

func (productMetric *ProductMetric) FeedDeleteQuery() string {
	return `
		DELETE FROM product_metrics
		WHERE id = $1
	`
}