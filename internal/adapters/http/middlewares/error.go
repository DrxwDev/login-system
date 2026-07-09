package middlewares

import "errors"

var (
	ErrMissingAuthHeader = errors.New("authentication header is missing")
	ErrInvalidAuthFormat = errors.New("invalid authentication header format")
	ErrInvalidToken      = errors.New("invalid access_token")
)
