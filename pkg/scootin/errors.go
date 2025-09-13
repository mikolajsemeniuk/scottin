package scootin

import "errors"

var (
	ErrCannotParseQuery = errors.New(`cannot parse query parameter`)
	ErrStatusInvalid    = errors.New("status can only be free or occupied")
	ErrScooterNotFound  = errors.New("scooter not found")
	ErrScooterOccupied  = errors.New("scooter is occupied")
	ErrReleaseScooter   = errors.New("scooter is not occupied")
)
