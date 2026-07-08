package user

import "strings"

type Password string

func NewPassword(value string) (Password, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return "", ErrPasswordRequired
	}

	if len(value) < 8 {
		return "", ErrPasswordTooShort
	}

	if len(value) > 255 {
		return "", ErrPasswordTooLong
	}

	return Password(value), nil
}

func (p Password) String() string {
	return string(p)
}
