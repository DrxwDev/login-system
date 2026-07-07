package user

import "errors"

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrPasswordTooLong    = errors.New("password too long")
	ErrUserIDRequired     = errors.New("user id is required")
	ErrNameRequired       = errors.New("name is required")
)
