package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type config struct {
	DSN           string `default:"postgres://wishlist:P@ssw0rd@localhost:5432/wishlist?sslmode=disable" envconfig:"DATABASE_URL"`
	MigrationsDir string `default:"migrations"                                                           envconfig:"MIGRATIONS_DIR"`
	Version       uint   `default:"1"                                                                    envconfig:"VERSION"`
}

//go:embed seed.sql
var seed string

func main() {
	var cfg config
	if err := envconfig.Process("MIGRATOR", &cfg); err != nil {
		log.Printf("failed to process config: %v", err)
		return
	}

	storage, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		log.Printf("failed to open database connection: %v", err)
		return
	}
	defer storage.Close()

	if _, err := storage.ExecContext(context.Background(), seed); err != nil {
		log.Printf("failed to execute seed SQL: %v", err)
		return
	}

	log.Println("seed data inserted successfully!")
}
