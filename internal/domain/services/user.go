package services

import "github.com/ahmetilboga2004/internal/domain/interfaces"

type userService struct {
	userRepo interfaces.IUserRepository
}

func NewUserService(userRepo interfaces.IUserRepository) *userService {
	return &userService{userRepo: userRepo}
}
