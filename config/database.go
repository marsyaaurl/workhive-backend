package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("parse DSN error:", err)
	}
	// optional: tune pool
	cfg.MaxConns = 5
	cfg.MinConns = 0
	cfg.MaxConnLifetime = time.Hour

	DB, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal("connect pool error:", err)
	}

	if err := DB.Ping(context.Background()); err != nil {
		log.Fatal("db ping error:", err)
	}
	log.Println("✅ PostgreSQL connected")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
