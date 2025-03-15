package order

import "github.com/jmoiron/sqlx"

func (order *Order) FeedGetID() *int {
	return &order.ID
}

func (order *Order) FeedCreateQuery() string {
	return `
		INSERT INTO orders (customer_id, total_price, status, created_at, updated_at)
		VALUES (:customer_id, :total_price, :status, :created_at, :updated_at)
		RETURNING id
	`
}

func (order *Order) FeedGetByIDQuery() string {
	return `
		SELECT id, customer_id, total_price, status, created_at, updated_at
		FROM orders
		WHERE id = $1
	`
}

func (order *Order) FeedGetAllQuery() string {
	return `
		SELECT id, customer_id, total_price, status, created_at, updated_at
		FROM orders
	`
}

func (order *Order) FeedUpdateDetailsQuery() string {
	return `
		UPDATE orders
		SET customer_id = :customer_id, 
			total_price = :total_price,
			status = :status
		WHERE id = :id
		RETURNING id, customer_id, total_price, status, created_at, updated_at
	`
}

func (order *Order) FeedDeleteQuery() string {
	return `
		DELETE FROM orders
		WHERE id = $1
	`
}

func (order_item *OrderItem) FeedGetID() *int {
	return &order_item.ID
}

func (order_item *OrderItem) FeedCreateQuery() string {
	return `
		INSERT INTO order_items (product_id, order_id, price, quantity, created_at, updated_at)
		VALUES (:product_id, :order_id, :price, :quantity, :created_at, :updated_at)
		RETURNING id
	`
}

func (order_item *OrderItem) FeedGetByIDQuery() string {
	return `
		SELECT id, product_id, order_id, price, quantity, created_at, updated_at
		FROM order_items
		WHERE id = $1
	`
}

func (order_item *OrderItem) FeedGetAllQuery() string {
	return `
		SELECT id, product_id, order_id, price, quantity, created_at, updated_at
		FROM order_items
	`
}

func (order_item *OrderItem) FeedUpdateDetailsQuery() string {
	return `
		UPDATE order_items
		SET product_id = :product_id, 
			order_id = :order_id, 
			price = :price,
			quantity = :quantity
		WHERE id = :id
		RETURNING id, product_id, order_id, price, quantity, created_at, updated_at
	`
}

func (order_item *OrderItem) FeedDeleteQuery() string {
	return `
		DELETE FROM order_items
		WHERE id = $1
	`
}

func GetOrdersByCustomerIDQuery(db *sqlx.DB, customer_id int, orders *[]Order) error {
	query := `	
		SELECT id, customer_id, total_price, status, created_at, updated_at
		FROM orders 
		WHERE customer_id = $1
	`
	return db.Select(orders, query, customer_id)
}

func GetOrderItemsByOrderIDQuery(db *sqlx.DB, order_id int, order_items *[]OrderItem) error {
	query := `	
		SELECT id, product_id, order_id, price, quantity, created_at, updated_at
		FROM order_items 
		WHERE order_id = $1
	`
	return db.Select(order_items, query, order_id)
}