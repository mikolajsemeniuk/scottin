package scootin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Status string

func (s Status) Valid() error {
	if s != "free" && s != "occupied" {
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
	query := r.URL.Query()

	lat1, err := strconv.ParseFloat(query.Get("latitude1"), 64)
	if err != nil {
		return FindScootersInput{}, ErrCannotParseQuery
	}

	lat2, err := strconv.ParseFloat(query.Get("latitude2"), 64)
	if err != nil {
		return FindScootersInput{}, ErrCannotParseQuery
	}

	lng1, err := strconv.ParseFloat(query.Get("longitude1"), 64)
	if err != nil {
		return FindScootersInput{}, ErrCannotParseQuery
	}

	lng2, err := strconv.ParseFloat(query.Get("longitude2"), 64)
	if err != nil {
		return FindScootersInput{}, ErrCannotParseQuery
	}

	out := FindScootersInput{
		Latitude1:  lat1,
		Latitude2:  lat2,
		Longitude1: lng1,
		Longitude2: lng2,
	}

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
}

func NewUpdateScooterInput(r *http.Request) (UpdateScooterInput, error) {
	var out UpdateScooterInput
	ID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return out, ErrCannotParseQuery
	}

	if err := json.NewDecoder(r.Body).Decode(&out); err != nil {
		return out, err
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
	Updated   time.Time `json:"updated"`
	Created   time.Time `json:"created"`
}
