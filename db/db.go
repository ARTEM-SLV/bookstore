package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"bookstore/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func InitDB(cfg *config.Config) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	var err error
	pool, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Connected to database")
}

func GetDB() *pgxpool.Pool {
	return pool
}
