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
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *userRepository) Update(user *models.User) error {
	result := r.db.Model(&user).Updates(&user)
	return result.Error
}

func (r *userRepository) ChangePassword(id uint, newPassword string) error {
	result := r.db.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword)
	return result.Error
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
