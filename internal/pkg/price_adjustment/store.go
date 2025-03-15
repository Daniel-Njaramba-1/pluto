package price_adjustment

func (price_adjustment *PriceAdjustment) FeedGetID() *int {
	return &price_adjustment.ID
}

func (price_adjustment *PriceAdjustment) FeedCreateQuery() string {
	return `
		INSERT INTO price_adjustments (product_id, old_price, new_price)
		VALUES (:product_id, :old_price, :new_price)
		RETURNING id
	`
}

func (price_adjustment *PriceAdjustment) FeedGetByIDQuery() string {
	return `
		SELECT id, product_id, old_price, new_price, created_at, updated_at
		FROM price_adjustments
		WHERE id = $1
	`
}

func (price_adjustment *PriceAdjustment) FeedGetAllQuery() string {
	return `
		SELECT id, product_id, old_price, new_price, created_at, updated_at
		FROM price_adjustments
	`
}

func (price_adjustment *PriceAdjustment) FeedUpdateDetailsQuery() string {
	return `
		UPDATE price_adjustments
		SET product_id = :product_id, 
			old_price = :old_price,
			new_price = :new_price
		WHERE id = :id
		RETURNING id, product_id, old_price, new_price, created_at, updated_at
	`
}

func (price_adjustment *PriceAdjustment) FeedDeleteQuery() string {
	return `
		DELETE FROM price_adjustments
		WHERE id = $1
	`
}