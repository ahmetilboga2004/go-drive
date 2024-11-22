package services

import (
	"errors"

	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
	"github.com/ahmetilboga2004/internal/infrastructure/utils/auth"
)

type userService struct {
	userRepo   interfaces.IUserRepository
	jwtService interfaces.IJwtService
}

func NewUserService(userRepo interfaces.IUserRepository, jwtService interfaces.IJwtService) interfaces.IUserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Register(user *models.User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Login(usernameOrEmail, password string) (string, string, error) {
	var user *models.User
	var err error
	user, err = s.userRepo.GetByUsername(usernameOrEmail)
	if err != nil {
		user, err = s.userRepo.GetByEmail(usernameOrEmail)
		if err != nil {
			return "", "", errors.New("invalid account")
		}
	}

	passCheck := auth.ComparePasswordHash(password, user.Password)
	if !passCheck {
		return "", "", errors.New("invalid password")
	}

	claims := map[string]any{
		"username": user.Username,
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, claims)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *userService) RefreshToken(refreshToken string) (string, error) {
	validatedToken, err := s.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	tokenClaims, err := s.jwtService.GetTokenClaims(validatedToken)
	if err != nil {
		return "", err
	}

	userIdFloat, ok := tokenClaims["user_id"].(float64)
	if !ok {
		return "", errors.New("invalid token payload: user_id not found or incorrect type")
	}

	userId := uint(userIdFloat)

	user, err := s.GetByID(userId)
	if err != nil {
		return "", err
	}

	newAccessTokenClaims := map[string]any{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"username":  user.Username,
	}

	newAccessToken, err := s.jwtService.GenerateAccessToken(userId, newAccessTokenClaims)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return &models.User{}, nil
}
