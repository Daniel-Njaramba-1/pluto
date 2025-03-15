package wishlist

func (wishlist *Wishlist) FeedGetID() *int {
	return &wishlist.ID
}

func (wishlist *Wishlist) FeedCreateQuery() string {
	return `
		INSERT INTO wishlists (customer_id, is_active)
		VALUES (:customer_id, :is_active)
		RETURNING id
	`
}

func (wishlist *Wishlist) FeedGetByIDQuery() string {
	return `
		SELECT id, customer_id, is_active, created_at, updated_at
		FROM wishlists
		WHERE id = $1
	`
}

func (wishlist *Wishlist) FeedGetAllQuery() string {
	return `
		SELECT id, customer_id, is_active, created_at, updated_at
		FROM wishlists
	`
}

func (wishlist *Wishlist) FeedUpdateDetailsQuery() string {
	return `
		UPDATE wishlists
		SET customer_id = :customer_id, 
			is_active = :is_active
		WHERE id = :id
		RETURNING id, customer_id, is_active, created_at, updated_at
	`
}

func (wishlist *Wishlist) FeedDeleteQuery() string {
	return `
		DELETE FROM wishlists
		WHERE id = $1
	`
}

func (wishlist_item *WishlistItem) FeedGetID() *int {
	return &wishlist_item.ID
}

func (wishlist_item *WishlistItem) FeedCreateQuery() string {
	return `
		INSERT INTO wishlist_items (wishlist_id, product_id)
		VALUES (:wishlist_id, :product_id)
		RETURNING id
	`
}

func (wishlist_item *WishlistItem) FeedGetByIDQuery() string {
	return `
		SELECT id, wishlist_id, product_id, created_at, updated_at
		FROM wishlist_items
		WHERE id = $1
	`
}

func (wishlist_item *WishlistItem) FeedGetAllQuery() string {
	return `
		SELECT id, wishlist_id, product_id, created_at, updated_at
		FROM wishlist_items
	`
}

func (wishlist_item *WishlistItem) FeedUpdateDetailsQuery() string {
	return `
		UPDATE wishlist_items
		SET wishlist_id = :wishlist_id, 
			product_id = :product_id
		WHERE id = :id
		RETURNING id, wishlist_id, product_id, created_at, updated_at
	`
}

func (wishlist_item *WishlistItem) FeedDeleteQuery() string {
	return `
		DELETE FROM wishlist_items
		WHERE id = $1
	`
}