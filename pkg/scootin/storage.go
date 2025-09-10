package scootin

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	store *sql.DB
}

func NewPostgresStorage(ctx context.Context, source string) (*PostgresStorage, error) {
	store, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := store.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	store.SetMaxOpenConns(25)
	store.SetMaxIdleConns(5)
	store.SetConnMaxLifetime(5 * time.Minute)

	return &PostgresStorage{store: store}, nil
}

func (s *PostgresStorage) Close() error {
	return s.store.Close()
}

func (s *PostgresStorage) FindScooters(ctx context.Context, lat1, lng1, lat2, lng2 float64, status string) ([]Scooter, error) {
	query := `
		SELECT id, city, status, COALESCE(latitude, 0), COALESCE(longitude, 0), updated, created
		FROM scooters
		WHERE latitude BETWEEN $1 AND $2
		AND longitude BETWEEN $3 AND $4
	`

	// Make code not order sensitive.
	args := []any{
		min(lat1, lat2), max(lat1, lat2),
		min(lng1, lng2), max(lng1, lng2),
	}

	if status != "" {
		query += " AND status = $5"
		args = append(args, string(status))
	}

	rows, err := s.store.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error finding scooters: %w", err)
	}
	defer rows.Close()

	// Make scooters initialized slice to return empty array instead of null
	scooters := []Scooter{}
	for rows.Next() {
		var scooter Scooter
		dst := []any{&scooter.ID, &scooter.City, &scooter.Status, &scooter.Latitude, &scooter.Longitude, &scooter.Updated, &scooter.Created}
		if err := rows.Scan(dst...); err != nil {
			return nil, fmt.Errorf("error scanning scooter row: %w", err)
		}

		scooters = append(scooters, scooter)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating scooter rows: %w", err)
	}

	if len(scooters) == 0 {
		return nil, ErrNoScootersFound
	}

	return scooters, nil
}

func (s *PostgresStorage) UpdateScooter(ctx context.Context, scooter Scooter) error {
	query := `
		UPDATE scooters
		SET status = $2,
		    latitude = $3,
		    longitude = $4,
		    updated = $5
		WHERE id = $1
	`

	args := []any{
		scooter.ID,
		scooter.Status,
		scooter.Latitude,
		scooter.Longitude,
		time.Now().UTC(),
	}

	result, err := s.store.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating scooter: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if affected == 0 {
		return ErrScooterNotFound
	}

	return nil
}
