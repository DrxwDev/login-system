package user

import (
	"strings"
)

type UserID string

func NewUserID(value string) (UserID, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return "", ErrUserIDRequired
	}

	return UserID(value), nil
}

func (id UserID) String() string {
	return string(id)
}
