package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DrxwDev/login-system/internal/platform/config"
)

func NewPool(cfg config.DatabaseConfig) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		log.Fatalf("unable to parse database config: %v", err)
	}

	// Optional pool settings
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	return pool
}
