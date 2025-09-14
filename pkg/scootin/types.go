package scootin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusFree     = Status("free")
	StatusOccupied = Status("occupied")
)

func (s Status) Valid() error {
	if s != StatusFree && s != StatusOccupied {
		return ErrStatusInvalid
	}

	return nil
}

type FindScootersInput struct {
	Latitude1  float64
	Latitude2  float64
	Longitude1 float64
	Longitude2 float64
	Status     Status
}

func NewFindScootersInput(r *http.Request) (FindScootersInput, error) {
	var out FindScootersInput
	query := r.URL.Query()

	lat1, err := strconv.ParseFloat(query.Get("latitude1"), 64)
	if err != nil {
		return out, ErrCannotParseQuery
	}

	lat2, err := strconv.ParseFloat(query.Get("latitude2"), 64)
	if err != nil {
		return out, ErrCannotParseQuery
	}

	lng1, err := strconv.ParseFloat(query.Get("longitude1"), 64)
	if err != nil {
		return out, ErrCannotParseQuery
	}

	lng2, err := strconv.ParseFloat(query.Get("longitude2"), 64)
	if err != nil {
		return out, ErrCannotParseQuery
	}

	out.Latitude1 = lat1
	out.Latitude2 = lat2
	out.Longitude1 = lng1
	out.Longitude2 = lng2

	status := Status(query.Get("status"))
	if status == "" {
		return out, nil
	}

	if err := status.Valid(); err != nil {
		return out, ErrStatusInvalid
	}

	out.Status = status

	return out, nil
}

type UpdateScooterInput struct {
	ID        uuid.UUID `json:"id"`
	Status    Status    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ClientID  string    `json:"client_id"`
}

func NewUpdateScooterInput(r *http.Request) (UpdateScooterInput, error) {
	var out UpdateScooterInput
	ID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return out, ErrCannotParseQuery
	}

	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		return out, fmt.Errorf("failed to decode request body: %w", err)
	}

	if err := out.Status.Valid(); err != nil {
		return out, ErrStatusInvalid
	}

	out.ID = ID

	return out, nil
}

type Scooter struct {
	ID        uuid.UUID `json:"id"`
	City      string    `json:"city"`
	Status    string    `json:"status"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	ClientID  string    `json:"client_id"`
	Updated   time.Time `json:"updated"`
	Created   time.Time `json:"created"`
}
