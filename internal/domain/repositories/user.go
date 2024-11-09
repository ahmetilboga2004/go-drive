package repositories

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]*models.User, error) {
	return []*models.User{}, nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	return &models.User{}, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *userRepository) Create(user *models.User) error {
	return nil
}

func (r *userRepository) Update(id uint, user *models.User) error {
	return nil
}

func (r *userRepository) ChangePassword(id uint, newPassword string) error {
	return nil
}

func (r *userRepository) Delete(id uint) error {
	return nil
}
