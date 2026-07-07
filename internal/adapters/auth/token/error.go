package token

import "errors"

var (
	ErrUserIDNotProvided   = errors.New("user id is missing")
	ErrUnableToCreateToken = errors.New("unable to create token")
	ErrInvalidTokenMethod  = errors.New("invalid token method")
	ErrUnableToParseToken  = errors.New("unable to parse token")
	ErrInvalidAccessToken  = errors.New("invalid access token")
)
