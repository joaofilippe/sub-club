package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/subclub/internal/config"
)

var connection *Connection

type Connection struct {
	appConfig *config.Config
	db *sqlx.DB
}

func NewConnection(ctx context.Context, appConfig *config.Config) (*Connection, error) {
	db, err := sqlx.ConnectContext(ctx, appConfig.DatabaseDriver, appConfig.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	connection = &Connection{
		appConfig: appConfig,
		db: db,
	}

	return connection, nil
}

func (c *Connection) GetDB() *sqlx.DB {
	return c.db
}
