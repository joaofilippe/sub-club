package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewConnection(ctx context.Context, connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
