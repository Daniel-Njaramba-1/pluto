package review

import "github.com/jmoiron/sqlx"

func (review *Review) FeedGetID() *int {
	return &review.ID
}

func (review *Review) FeedCreateQuery() string {
	return `
		INSERT INTO reviews (customer_id, product_id, rating, review_text, created_at, updated_at)
		VALUES (:customer_id, :product_id, :rating, :review_text, :created_at, :updated_at)
		RETURNING id
	`
}

func (review *Review) FeedGetByIDQuery() string {
	return `
		SELECT id, customer_id, product_id, rating, review_text, created_at, updated_at
		FROM reviews
		WHERE id = $1
	`
}

func (review *Review) FeedGetAllQuery() string {
	return `
		SELECT id, customer_id, product_id, rating, review_text, created_at, updated_at
		FROM reviews
	`
}

func (review *Review) FeedUpdateDetailsQuery() string {
	return `
		UPDATE reviews
		SET customer_id = :customer_id, 
			product_id = :product_id,
			rating = :rating,
			review_text = :review_text
		WHERE id = :id
		RETURNING id, customer_id, product_id, rating, review_text, created_at, updated_at
	`
}

func (review *Review) FeedDeleteQuery() string {
	return `
		DELETE FROM reviews
		WHERE id = $1
	`
}

func GetReviewsByCustomerIDQuery(db *sqlx.DB, customer_id int, reviews *[]Review) error {
	query := `	
		SELECT id, customer_id, product_id, rating, review_text, created_at, updated_at
		FROM reviews 
		WHERE customer_id = $1
	`
	return db.Select(reviews, query, customer_id)
}