package interfaces

import "github.com/ahmetilboga2004/internal/domain/models"

type IFileRepository interface {
	IBaseRepository[models.File]
	// ChangeFilePath(fileID uint, newPath string) error
	// CreateShareLink(fileID uint) (string, error)
}

type IFileService interface {
	GetAll() ([]*models.File, error)
	GetByID(id uint) (*models.File, error)
	Create(file *models.File) error
	Update(id uint, file *models.File) error
	Delete(id uint) error
}
