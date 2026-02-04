// Package user
package user

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid_email")
	ErrUserExists   = errors.New("user_exists")
)
