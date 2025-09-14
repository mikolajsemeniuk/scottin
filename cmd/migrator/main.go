package main

import (
	"embed"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/kelseyhightower/envconfig"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type config struct {
	DSN           string `default:"postgres://wishlist:P@ssw0rd@localhost:5432/wishlist?sslmode=disable" envconfig:"DATABASE_URL"`
	MigrationsDir string `default:"migrations"                                                           envconfig:"MIGRATIONS_DIR"`
	Version       uint   `default:"1"                                                                    envconfig:"VERSION"`
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

	if err := instance.Migrate(cfg.Version); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Failed to migrate to version %d: %v", cfg.Version, err)
		return
	}
}
