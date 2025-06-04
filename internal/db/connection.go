package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/redis/go-redis/v9"
)

type Connections struct {
	Postgres *pg.DB
	Redis    *redis.Client
}

func NewConnections(ctx context.Context) (*Connections, error) {
	// Setup PostgreSQL
	pgOptions := &pg.Options{
		Addr:     fmt.Sprintf("%s:%s", getEnv("POSTGRES_HOST", "localhost"), getEnv("POSTGRES_PORT", "5432")),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "victoria7"),
		Database: getEnv("POSTGRES_DB", "meetup_dev"),
	}
	db := pg.Connect(pgOptions)
	if _, err := db.Exec("SELECT 1"); err != nil {
		return nil, nilConnectionError("PostgreSQL", err)
	}

	// Setup Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "6379")),
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, nilConnectionError("Redis", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL and Redis")

	return &Connections{
		Postgres: db,
		Redis:    rdb,
	}, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func nilConnectionError(service string, err error) error {
	log.Printf("❌ %s connection failed: %v", service, err)
	return fmt.Errorf("%s connection error: %w", service, err)
}
