package scootin

import "errors"

var (
	ErrCannotParseQuery              = errors.New(`cannot parse query parameter`)
	ErrStatusInvalid                 = errors.New("status can only be free or occupied")
	ErrScooterNotFound               = errors.New("scooter not found")
	ErrClientIDRequired              = errors.New("client_id is required when occupying scooter")
	ErrOnlyOccupyingClientCanRelease = errors.New("only the occupying client can release the scooter")
	ErrOnlyOccupyingClientCanUpdate  = errors.New("only the occupying client can update scooter position")
	ErrScooterAlreadyOccupied        = errors.New("scooter is already occupied by another client")
)
