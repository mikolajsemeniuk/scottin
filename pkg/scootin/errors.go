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
	ErrScooterStatusChanged          = errors.New("scooter status was changed by another client")
	ErrAPIRequestFailed              = errors.New("API request failed")
	ErrScooterAlreadyReserved        = errors.New("scooter already reserved")
	ErrNoFreeScooters                = errors.New("no free scooters available")
	ErrReservationFailed             = errors.New("failed to reserve scooter")
)
