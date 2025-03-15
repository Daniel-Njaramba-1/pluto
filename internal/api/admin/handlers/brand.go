package admin_handlers

import (
	"net/http"
	"pluto/internal/pkg/brand"

	"github.com/labstack/echo/v4"
)

type BrandHandler struct {
	brandService *brand.BrandService
}

func NewBrandHandler(brandService *brand.BrandService) *BrandHandler {
	return &BrandHandler{brandService: brandService}
}

func (h *BrandHandler) Create(c echo.Context) error {
	brand := new(brand.Brand)
	if err := c.Bind(brand); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	id, err := h.brandService.Create(brand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
		"message": "Brand created successfully",
	})
}
