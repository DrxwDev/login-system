// Package handlers
package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/DrxwDev/login-system/internal/adapters/http/dto"
	"github.com/DrxwDev/login-system/internal/adapters/http/mapper"
	"github.com/DrxwDev/login-system/internal/adapters/http/middlewares"
	usecase "github.com/DrxwDev/login-system/internal/core/usecase/user"
)

type UserHandler struct {
	register *usecase.RegisterUseCase
	validate *validator.Validate
	login    *usecase.LoginUseCase
	find     *usecase.GetUserByIDUseCase
}

func NewUserHandler(register *usecase.RegisterUseCase, validate *validator.Validate, login *usecase.LoginUseCase, find *usecase.GetUserByIDUseCase) *UserHandler {
	return &UserHandler{
		register: register,
		validate: validate,
		login:    login,
		find:     find,
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
			"message": "invalid forms fields",
			"error":   err.Error(),
		})
		return
	}

	params := mapper.UserRegisterParams(payload)
	newUser, token, err := h.register.Register(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"status":  "failed",
			"message": "user already exists",
			"error":   err.Error(),
		})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Path:     "/",
		Value:    token,
		MaxAge:   3600,
		Expires:  time.Now().Add(time.Hour),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	ctx.JSON(http.StatusCreated, mapper.UserDomainToDTO(newUser))
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var payload dto.LoginRequest
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
			"message": "invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	params := mapper.UserLoginParams(payload)
	user, token, err := h.login.Login(ctx.Request.Context(), params)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Path:     "/",
		Value:    token,
		MaxAge:   3600,
		Expires:  time.Now().Add(time.Hour),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   mapper.UserDomainToDTO(user),
	})
}

func (h *UserHandler) Logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access_token",
		Path:     "/",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Secure:   false,
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	})

	ctx.Status(http.StatusOK)
}

func (h *UserHandler) User(ctx *gin.Context) {
	claims, ok := middlewares.GetClaims(ctx)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	user, err := h.find.GetUser(ctx.Request.Context(), claims.Subject)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   mapper.UserDomainToDTO(user),
	})
}
