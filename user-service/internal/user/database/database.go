package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("Database url is empty")
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	if er := pool.Ping(ctx); er != nil {
		return nil, fmt.Errorf("Unable to ping database: %v", er)
	}

	sql, err := os.ReadFile("user_database.sql")
	if err != nil {
		return nil, fmt.Errorf("Unable to read user_database.sql: %v", err)
	}
	if _, err := pool.Exec(ctx, string(sql)); err != nil {
		return nil, fmt.Errorf("Unable to initialize database schema: %v", err)
	}
	return pool, nil
}
