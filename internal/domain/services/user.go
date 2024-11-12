package services

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
)

type userService struct {
	userRepo interfaces.IUserRepository
}

func NewUserService(userRepo interfaces.IUserRepository) interfaces.IUserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(user *models.User) error {
	return nil
}

func (s *userService) Login(usernameOrEmail, password string) (string, error) {
	return "token", nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return &models.User{}, nil
}
