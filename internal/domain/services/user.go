package services

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
	"github.com/ahmetilboga2004/internal/infrastructure/utils/auth"
)

type userService struct {
	userRepo interfaces.IUserRepository
}

func NewUserService(userRepo interfaces.IUserRepository) interfaces.IUserService {
	return &userService{userRepo: userRepo}
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

func (s *userService) Login(usernameOrEmail, password string) (string, error) {
	return "token", nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return &models.User{}, nil
}
