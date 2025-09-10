package scootin

import "errors"

var (
	ErrCannotParseQuery = errors.New(`cannot parse query parameter`)
	ErrStatusInvalid    = errors.New("status can only be free or occupied")
	ErrNoScootersFound  = errors.New("no scooters found")
	ErrScooterNotFound  = errors.New("scooter not found")
)
