package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the SubClub API server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func startServer() {
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

	srv := server.NewServer()

	log.Fatal(srv.Start(":8080"))
}
