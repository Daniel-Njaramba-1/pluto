package admin_handlers

import (
	"net/http"
	"pluto/internal/pkg/admin"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminService *admin.AdminService
}

func NewAdminHandler (adminService *admin.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) Register (c echo.Context) error {
	admin := new(admin.Admin)
	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": "Invalid request body",
		})
	}

	token, err := h.adminService.Register(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]string {
		"token": token,
		"message": "Admin registered successfully",
	})
}

func (h *AdminHandler) Login (c echo.Context) error {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"error": "Invalid request body",
		})
	}

	token, err := h.adminService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string {
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string {
		"token": token,
		"message": "Login successful",
	})
}