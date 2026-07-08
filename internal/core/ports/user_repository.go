package ports

import (
	"context"

	"github.com/DrxwDev/login-system/internal/core/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, u user.User) error
	FindByEmail(ctx context.Context, email user.Email) (user.User, error)
	FindByID(ctx context.Context, id user.UserID) (user.User, error)
}
