package main

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/kelseyhightower/envconfig"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type config struct {
	DSN           string `envconfig:"DATABASE_URL"   default:"postgres://wishlist:P@ssw0rd@localhost:5432/wishlist?sslmode=disable"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"migrations"`
	Version       uint   `envconfig:"VERSION"        default:"1"`
}

func main() {
	var cfg config
	if err := envconfig.Process("MIGRATOR", &cfg); err != nil {
		log.Printf("Failed to process config: %v", err)
		return
	}

	driver, err := iofs.New(migrationsFS, cfg.MigrationsDir)
	if err != nil {
		log.Printf("Failed to create migration source: %v", err)
		return
	}

	instance, err := migrate.NewWithSourceInstance("iofs", driver, cfg.DSN)
	if err != nil {
		log.Printf("Failed to create migrate instance: %v", err)
		return
	}

	defer func() {
		if _, err := instance.Close(); err != nil {
			log.Printf("Failed to close migrate instance: %v", err)
			return
		}

		log.Printf("Successfully migrated to version %d\n", cfg.Version)
	}()

	if err := instance.Migrate(cfg.Version); err != nil && err != migrate.ErrNoChange {
		log.Printf("Failed to migrate to version %d: %v", cfg.Version, err)
		return
	}
}
