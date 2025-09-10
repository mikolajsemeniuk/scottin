package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/mikolajsemeniuk/nord/pkg/scootin"
)

type config struct {
	Name         string        `default:"scotting"                                                             envconfig:"NAME"`
	Listen       string        `default:":8080"                                                                envconfig:"LISTEN"`
	LogLevel     int           `default:"0"                                                                    envconfig:"LOG_LEVEL"`
	ReadTimeout  time.Duration `default:"5s"                                                                   envconfig:"WRITE_TIMEOUT"`
	WriteTimeout time.Duration `default:"30s"                                                                  envconfig:"IDLE_TIMEOUT"`
	IdleTimeout  time.Duration `default:"30s"                                                                  envconfig:"IDLE_TIMEOUT"`
	DatabaseURL  string        `default:"postgres://wishlist:P@ssw0rd@localhost:5432/wishlist?sslmode=disable" envconfig:"DATABASE_URL"`
}

func main() {
	ctx := context.Background()

	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Println("error parsing config: %w", err)
		return
	}

	storage, err := scootin.NewPostgresStorage(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer storage.Close()

	handler, err := scootin.NewHTTPHandler(storage)
	if err != nil {
		log.Println("Error creating handler:", err)
		return
	}

	router := http.NewServeMux()
	router.Handle("/", handler)

	server := &http.Server{
		Addr:         cfg.Listen,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Make sure the server is shutdown gracefully, specialy important in K8S environment.
	go func() {
		log.Printf("Starting server on %s", cfg.Listen)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
	}
}
