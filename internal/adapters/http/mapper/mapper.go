// Package mapper
package mapper

import (
	"github.com/DrxwDev/login-system/internal/adapters/http/dto"
	"github.com/DrxwDev/login-system/internal/core/domain/user"
	usecase "github.com/DrxwDev/login-system/internal/core/usecase/user"
)

func UserDomainToDTO(user user.User) dto.UserDTO {
	return dto.UserDTO{
		ID:        string(user.ID),
		Name:      user.Name,
		Email:     string(user.Email),
		CreatedAt: user.CreatedAt.String(),
	}
}

func UserRegisterParams(payload dto.RegisterRequest) usecase.RegisterParams {
	return usecase.RegisterParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
	}
}

func UserLoginParams(payload dto.LoginRequest) usecase.LoginParams {
	return usecase.LoginParams{
		Email:    payload.Email,
		Password: payload.Password,
	}
}
