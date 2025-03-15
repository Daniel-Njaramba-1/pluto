package admin

import (
	"errors"
	"fmt"
	"pluto/internal/lib/generics"
	"pluto/internal/lib/hashing"

	"github.com/jmoiron/sqlx"
)

// AdminService provides all admin-related operations
type AdminService struct {
    db *sqlx.DB
}

// NewAdminService creates a new admin service
func NewAdminService(db *sqlx.DB) *AdminService {
    return &AdminService{db: db}
}

// Register registers a new admin
func (s *AdminService) Register(admin *Admin) (string, error) {
    //log data from echo, logger.go is at package logger
    if admin.Username == "" || admin.Email == "" || admin.Password == "" {
        return "", errors.New("username, email, and password cannot be empty")
    }

    // Hash the password
    hashedPassword, err := hashing.HashPassword(admin.Password)
    if err != nil {
        return "", err
    }
    admin.PasswordHash = hashedPassword
    admin.Password = "" // clear it

    // Create the admin in the database
    _, err = generics.CreateModel(s.db, admin)
    if err != nil {
        return "", err
    }

    // Create a token for the new admin
    token, err := CreateToken(admin.Username)
    if err != nil {
        return "", err
    }

    return token, nil
}

// Login logs in an admin
func (s *AdminService) Login(username, password string) (string, error) {
    // Retrieve the admin by username
    var admin Admin
    err := GetByNameQuery(s.db, username, &admin)
    if err != nil {
        return "", fmt.Errorf("authentication failed: %w", err)
    }

    // Compare the provided password with the stored hash
    if !hashing.VerifyPassword(password, admin.PasswordHash) {
        return "", fmt.Errorf("authentication failed: %w", err)
    }

    // Create a token for the admin
    token, err := CreateToken(username)
    if err != nil {
        return "", err
    }

    return token, nil
}

// EditProfile updates an admin's profile
func (s *AdminService) EditProfile(admin *Admin) error {
    if admin.ID == 0 {
        return errors.New("admin ID cannot be zero")
    }

    if admin.Username == "" || admin.Email == "" {
        return errors.New("username and email cannot be empty")
    }

    return generics.UpdateModelDetails(s.db, admin)
}