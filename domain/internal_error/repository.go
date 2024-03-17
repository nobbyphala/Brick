package internal_error

import "errors"

var (
	ErrNoRowsAffected = errors.New("error expected there row be affected but got none")
)
