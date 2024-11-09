package handlers

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService interfaces.IUserService
	validator   *validator.Validate
}

func NewUserHandler(userService interfaces.IUserService) *userHandler {
	return &userHandler{
		userService: userService,
		validator:   validator.New(),
	}
}
