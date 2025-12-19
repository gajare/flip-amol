package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is empty")
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	// run migration
	sql, err := os.ReadFile("database.sql")
	if err != nil {
		return nil, err
	}

	if _, err := pool.Exec(ctx, string(sql)); err != nil {
		return nil, err
	}

	return pool, nil
}
