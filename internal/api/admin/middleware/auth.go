package admin_middleware

import (
	"net/http"
	"pluto/internal/pkg/admin"
	"strings"

	"github.com/labstack/echo/v4"
)

//protect Admin Routes
func AdminAuth() echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			//Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string {
					"error" : "Authorization header required",
				})
			}

			//Check if it's a Bearer token
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string {
					"error": "Bearer token required",
				})
			}

			//Extract Token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			//Validate Token
			claims, err := admin.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
                    "error": "Invalid token: " + err.Error(),
                })
			}

			// Add claims to context
            c.Set("username", claims.Username)
            
            return next(c)
		}
	}
}