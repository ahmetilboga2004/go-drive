package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmetilboga2004/internal/application/dto"
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
	"github.com/ahmetilboga2004/internal/infrastructure/utils"
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

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, utils.InvalidData, err)
		return
	}
	if err := h.validator.Struct(userReq); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, utils.InvalidData, err)
		return
	}

	user := models.User{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Username:  userReq.Username,
		Email:     userReq.Email,
		Password:  userReq.Password,
	}

	err := h.userService.Register(&user)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, utils.InternalServerError, err)
		return
	}

	utils.SuccessResponse(w, http.StatusCreated, utils.UserRegisteredSuccessfully, nil)
}
