package admin

import (
	"github.com/jmoiron/sqlx"
)

func (admin *Admin) FeedGetID() *int {
	return &admin.ID
}

func (admin *Admin) FeedCreateQuery() string {
	return `
		INSERT INTO admins (username, email, password_hash, is_active)
		VALUES (:username, :email, :password_hash, :is_active)
		RETURNING id
	`
}

func (admin *Admin) FeedGetByIDQuery() string {
	return `
		SELECT id, username, email, is_active, created_at, updated_at
		FROM admins
		WHERE id = $1
	`
}

func (admin *Admin) FeedGetAllQuery() string {
	return `
		SELECT id, username, email, is_active, created_at, updated_at
		FROM admins
		ORDER BY id ASC
	`
}

func (admin *Admin) FeedUpdateDetailsQuery() string {
	return `
		UPDATE admins
		SET username = :username,
			email = :email,
		WHERE id = :id
		RETURNING id, username, email, is_active, created_at, updated_at
	`
}

func (admin *Admin) FeedDeactivateQuery() string {
	return `
		UPDATE admins
		SET is_active = FALSE
		WHERE id = $1
		RETURNING id, username, email, is_active, created_at, updated_at
	`
}

func (admin *Admin) FeedReactivateQuery() string {
	return `
		UPDATE admins
		SET is_active = TRUE
		WHERE id = $1
		RETURNING id, username, email, is_active, created_at, updated_at
	`
}

func (admin *Admin) FeedDeleteQuery() string {
	return `
		DELETE FROM admins
		WHERE id = $1
	`
}

func GetByNameQuery(db *sqlx.DB, name string, admin *Admin) error {
	query := 
	`	SELECT username, email, is_active, created_at, updated_at
		FROM admins 
		WHERE username = $1
	`
	return db.Get(admin, query, name)
}

func SearchByNameQuery(db *sqlx.DB, name string, admin *Admin) error {
	query := 
	`	SELECT username, email, is_active, created_at, updated_at
		FROM admins 
		WHERE username ILIKE $1
	`
	return db.Select(admin, query, name)
}

// for comparing with password attempts on login
func GetPasswordQuery(db *sqlx.DB, id int, admin *Admin) error {
	query := 
	`	SELECT password_hash
		FROM admins 
		WHERE username = $1
	`
	return db.Get(admin, query, id)
}

// update password using ID
func UpdatePasswordQuery(db *sqlx.DB, id int, admin *Admin) error {
	query := 
	`	UPDATE admins
		SET password_hash = $1
		WHERE id = $2
	`
	_, err := db.NamedExec(query, admin)
	return err
}




