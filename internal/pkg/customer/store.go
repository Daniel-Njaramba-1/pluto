package customer

import "github.com/jmoiron/sqlx"

func (customer *Customer) FeedGetID() *int {
	return &customer.ID
}

func (customer *Customer) FeedCreateQuery() string {
	return `
		INSERT INTO customers (firstname, lastname, username, email, password_hash, phone, address, is_active)
		VALUES (:firstname, :lastname, :username, :email, :password_hash, :phone, :address, :is_active)
		RETURNING id
	`
}

func (customer *Customer) FeedGetByIDQuery() string {
	return `
		SELECT id, firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
		FROM customers
		WHERE id = $1
	`
}

func (customer *Customer) FeedGetAllQuery() string {
	return `
		SELECT id, firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
		FROM customers
	`
}

func (customer *Customer) FeedUpdateDetailsQuery() string {
	return `
		UPDATE customers
		SET firstname = :firstname,
			lastname = :lastname,
			username = :username,
			email = :email,
			phone = :phone,
			address = :address
		WHERE id = :id
		RETURNING id, firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
	`
}

func (customer *Customer) FeedDeleteQuery() string {
	return `
		DELETE FROM customers
		WHERE id = $1
	`
}

func (customer *Customer) FeedDeactivateQuery() string {
	return `
		UPDATE customers
		SET is_active = FALSE
		WHERE id = $1
		RETURNING id, firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
	`
}

func (customer *Customer) FeedReactivateQuery() string {
	return `
		UPDATE customers
		SET is_active = TRUE
		WHERE id = $1
		RETURNING id, firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
	`
}

func GetByNameQuery(db *sqlx.DB, name string, customer *Customer) error {
	query :=
		`	SELECT firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
		FROM customers 
		WHERE username = $1
	`
	return db.Get(customer, query, name)
}

func SearchByNameQuery(db *sqlx.DB, name string, customer *Customer) error {
	query :=
		`	SELECT firstname, lastname, username, email, password_hash, phone, address, is_active, created_at, updated_at
		FROM customers 
		WHERE username ILIKE $1
	`
	return db.Select(customer, query, name)
}

// for comparing with password attempts on login
func GetPasswordQuery(db *sqlx.DB, id int, customer *Customer) error {
	query :=
		`	SELECT password_hash
		FROM customers 
		WHERE username = $1
	`
	return db.Get(customer, query, id)
}

// update password using ID
func UpdatePasswordQuery(db *sqlx.DB, id int, customer *Customer) error {
	query :=
		`	UPDATE customers
		SET password_hash = $1
		WHERE id = $2
	`
	_, err := db.NamedExec(query, customer)
	return err
}