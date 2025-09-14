package scootin

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		SELECT id, city, status, latitude, longitude, client_id, updated, created
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
		dst := []any{&scooter.ID, &scooter.City, &scooter.Status, &scooter.Latitude, &scooter.Longitude, &scooter.ClientID, &scooter.Updated, &scooter.Created}
		if err := rows.Scan(dst...); err != nil {
			return nil, fmt.Errorf("error scanning scooter row: %w", err)
		}

		scooters = append(scooters, scooter)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating scooter rows: %w", err)
	}

	return scooters, nil
}

func (s *PostgresStorage) FindScooter(ctx context.Context, id uuid.UUID) (*Scooter, error) {
	var scooter Scooter
	query := "SELECT id, city, status, latitude, longitude, client_id, updated, created FROM scooters WHERE id = $1"
	args := []any{&scooter.ID, &scooter.City, &scooter.Status, &scooter.Latitude, &scooter.Longitude, &scooter.ClientID, &scooter.Updated, &scooter.Created}
	err := s.store.QueryRowContext(ctx, query, id).Scan(args...)
	if err == sql.ErrNoRows {
		return nil, ErrScooterNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query scooter: %w", err)
	}

	return &scooter, nil
}

func (s *PostgresStorage) UpdateScooter(ctx context.Context, sc Scooter) error {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
		SELECT id, city, status, latitude, longitude, client_id, updated, created 
		FROM scooters
		WHERE id = $1
	`

	var current Scooter
	args := []any{
		&current.ID, &current.City, &current.Status, &current.Latitude,
		&current.Longitude, &current.ClientID, &current.Updated, &current.Created,
	}

	err = tx.QueryRowContext(ctx, query, sc.ID).Scan(args...)
	if err == sql.ErrNoRows {
		return ErrScooterNotFound
	}

	if err != nil {
		return fmt.Errorf("failed to query scooter: %w", err)
	}

	// Only allow occupation if scooter is free and client provides valid ID.
	if sc.Status == string(StatusOccupied) && current.Status == string(StatusFree) && sc.ClientID == "" {
		return ErrClientIDRequired
	}

	// Prevent occupation of already occupied scooter by different client
	if sc.Status == string(StatusOccupied) && current.Status == string(StatusOccupied) && sc.ClientID != current.ClientID {
		return ErrScooterAlreadyOccupied
	}

	// Only allow release if scooter is occupied and client owns it.
	if sc.Status == string(StatusFree) && current.Status == string(StatusOccupied) && sc.ClientID != current.ClientID {
		return ErrOnlyOccupyingClientCanRelease
	}

	query = `
		UPDATE scooters 
		SET status = $1, latitude = $2, longitude = $3, client_id = $4, updated = $5 
		WHERE id = $6
	`

	args = []any{
		sc.Status, sc.Latitude, sc.Longitude,
		sc.ClientID, sc.Updated, sc.ID,
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update scooter: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("scooter %s status was changed by another client", sc.ID)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
