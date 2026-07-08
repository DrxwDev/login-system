package user

import (
	"net/mail"
	"strings"
)

type Email string

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return "", ErrEmailRequired
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return "", ErrInvalidEmailFormat
	}

	return Email(value), nil
}

func (e Email) String() string {
	return string(e)
}
