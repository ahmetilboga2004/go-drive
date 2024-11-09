package interfaces

import "github.com/ahmetilboga2004/internal/domain/models"

type IFileRepository interface {
	IBaseRepository[models.File]
	// ChangeFilePath(fileID uint, newPath string) error
	// CreateShareLink(fileID uint) (string, error)
}
