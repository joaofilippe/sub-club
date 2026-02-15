package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/database"
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
}
