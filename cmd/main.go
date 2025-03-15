package main

import (
	"log"
	admin_routes "pluto/internal/api/admin/routes"
	"pluto/internal/config"
	"pluto/internal/lib/logger"
	"pluto/internal/server"
)

func main() {
    config.LoadEnv()

    logger.InitLogger()
    defer logger.CloseLogger()

    server, err := server.NewServer()
    if err != nil {
        log.Fatalf("Failed to initialize server: %v", err)
    }
    defer server.Close()

    admin_routes.SetupRoutes(server.Echo(), server.AdminHandler(), server.BrandHandler())

    port := config.GetEnv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    logger.LogInfo("Server started successfully")
    if err := server.Echo().Start(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
