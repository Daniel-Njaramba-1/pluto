package admin_routes

import (
	admin_handlers "pluto/internal/api/admin/handlers"
	"pluto/internal/api/admin/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, 
	adminHandler *admin_handlers.AdminHandler,
	brandHandler *admin_handlers.BrandHandler) {

	e.POST("/api/admin/register", adminHandler.Register)
	e.POST("/api/admin/login", adminHandler.Login)

	admin := e.Group("/api/admin")
	admin.Use(admin_middleware.AdminAuth())

	admin.POST("/brands", brandHandler.Create)
}