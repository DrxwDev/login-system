package user

import (
	"context"

	"github.com/DrxwDev/login-system/internal/core/domain/user"
	"github.com/DrxwDev/login-system/internal/core/ports"
)

type GetUserByIDUseCase struct {
	repo ports.UserRepository
}

func NewGetUserUseCase(repo ports.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		repo: repo,
	}
}

func (uc *GetUserByIDUseCase) GetUser(ctx context.Context, id string) (user.User, error) {
	if id == "" {
		return user.User{}, user.ErrInvalidUserID
	}

	u, err := uc.repo.FindByID(ctx, user.UserID(id))
	if err != nil {
		return user.User{}, ErrUserNotFound
	}

	return u, nil
}
