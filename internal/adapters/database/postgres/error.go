package postgres

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUnableToCreateUser = errors.New("unable to create user")
)
