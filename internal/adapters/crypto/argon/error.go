package argon

import "errors"

var (
	ErrPasswordRequired = errors.New("password is required")
	ErrHashRequired     = errors.New("hash is required")
	ErrPasswordMismatch = errors.New("password does not match")
	ErrHashMismatch     = errors.New("unable to verify hash")
)
