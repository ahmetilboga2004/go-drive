package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ahmetilboga2004/internal/application/dto"
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
	"github.com/ahmetilboga2004/internal/infrastructure/config"
	httphelper "github.com/ahmetilboga2004/internal/infrastructure/utils/httpHelper"
)

type userHandler struct {
	userService interfaces.IUserService
}

func NewUserHandler(userService interfaces.IUserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid data", err)
		return
	}
	if err := httphelper.Validator.Struct(userReq); err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid data", err)
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
		httphelper.ErrorResponse(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusCreated, "user register successfully", nil)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid data", err)
		return
	}

	if err := httphelper.Validator.Struct(userReq); err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid data", err)
		return
	}

	accessToken, refreshToken, err := h.userService.Login(userReq.UsernameOrEmail, userReq.Password)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(config.JWT.RefreshTokenExpiration),
	}

	http.SetCookie(w, cookie)
	httphelper.SuccessResponse(w, http.StatusOK, "Login successfully", map[string]any{
		"accessToken": accessToken,
	})

}

func (h *userHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "refreshToken cookie required", err)
		return
	}

	refreshToken := cookie.Value

	newAccessToken, err := h.userService.RefreshToken(refreshToken)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, "invalid refreshToken payload", err)
		return
	}

	httphelper.SuccessResponse(w, http.StatusOK, "refreshed token", map[string]any{
		"accessToken": newAccessToken,
	})
}
