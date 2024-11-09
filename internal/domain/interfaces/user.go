package interfaces

import "github.com/ahmetilboga2004/internal/domain/models"

type IUserRepository interface {
	IBaseRepository[models.User]
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	ChangePassword(userID uint, newPassword string) error
}
