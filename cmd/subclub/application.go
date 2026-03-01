package main

import (
	"context"
	"log"

	"github.com/joaofilippe/subclub/internal/application"
	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/server"
	"github.com/spf13/cobra"
)

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Start the SubClub application",
	Run: func(cmd *cobra.Command, args []string) {
		startApplication()
	},
}

func init() {
	rootCmd.AddCommand(applicationCmd)
}

func startApplication() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	ctx := context.Background()
	dbConnection, err := database.NewConnection(ctx, cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer dbConnection.Close()

	srv := server.NewServer()

	app := application.New(srv, dbConnection.GetDB())

	log.Fatal(app.Start(cfg.Port))
}
