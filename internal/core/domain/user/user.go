// Package user
package user

import (
	"strings"
	"time"
)

type User struct {
	ID        UserID
	Name      string
	Email     Email
	Password  Password
	CreatedAt time.Time
}

func NewUser(id UserID, name string, email Email, password Password) (User, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return User{}, ErrNameRequired
	}

	if len(name) < 5 {
		return User{}, ErrNameTooShort
	}

	if len(name) > 100 {
		return User{}, ErrNameTooLong
	}

	return User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
