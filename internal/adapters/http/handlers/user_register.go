// Package handlers
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/DrxwDev/login-system/internal/adapters/http/dto"
	"github.com/DrxwDev/login-system/internal/adapters/http/mapper"
	usecase "github.com/DrxwDev/login-system/internal/core/usecase/user"
)

type UserHandler struct {
	register *usecase.RegisterUseCase
	validate *validator.Validate
}

func NewUserHandler(register *usecase.RegisterUseCase, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		register: register,
		validate: validate,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var payload dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid form values",
			"error":   err.Error(),
		})
		return
	}

	if err := h.validate.Struct(payload); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "failed",
			"message": "unable to register",
			"error":   err.Error(),
		})
		return
	}

	params := mapper.UserRegisterParams(payload)
	newUser, err := h.register.Register(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"status":  "failed",
			"message": "unable to create user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, mapper.UserDomainToDTO(newUser))
}
