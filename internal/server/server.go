package server

import (
	admin_handlers "pluto/internal/api/admin/handlers"
	"pluto/internal/db"
	"pluto/internal/pkg/admin"
	"pluto/internal/pkg/brand"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	db       *sqlx.DB
	echo     *echo.Echo
	adminSvc *admin.AdminService
	brandSvc *brand.BrandService
	adminHdl *admin_handlers.AdminHandler
	brandHdl *admin_handlers.BrandHandler
}

func NewServer() (*Server, error) {
	db, err := db.ConnDB()
	if err != nil {
		return nil, err
	}

	e := echo.New()

	//enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, //admin dashboard on svelte 
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	adminSvc := admin.NewAdminService(db)
	brandSvc := brand.NewBrandService(db)

	server := &Server{
		db:       db,
		echo:     e,
		adminSvc: adminSvc,
		brandSvc: brandSvc,
		adminHdl: admin_handlers.NewAdminHandler(adminSvc),
		brandHdl: admin_handlers.NewBrandHandler(brandSvc),
	}

	return server, nil
}

func (s *Server) Close() {
    if s.db != nil {
        s.db.Close()
    }
}

func (s *Server) Echo() *echo.Echo {
    return s.echo
}

func (s *Server) AdminHandler() *admin_handlers.AdminHandler {
    return s.adminHdl
}

func (s *Server) BrandHandler() *admin_handlers.BrandHandler {
    return s.brandHdl
}


