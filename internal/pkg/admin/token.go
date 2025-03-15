package admin

import (
    "errors"
    "pluto/internal/config"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var adminKey = []byte(config.GetEnv("admin_key"))

// Claims represents the JWT claims
type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// CreateToken generates a JWT token for an admin
func CreateToken(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(adminKey)

    return tokenString, err
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return adminKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}