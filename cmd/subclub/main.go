package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/database"
	"github.com/joaofilippe/subclub/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Starting SubClub...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx := context.Background()
	dbConnection, err := database.NewConnection(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer dbConnection.Close()

	fmt.Println("Successfully connected to the database!")

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.ConfigureLogger())
	e.Use(echoMiddleware.Recover())

	// Routes
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
