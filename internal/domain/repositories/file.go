package repositories

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
	"gorm.io/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) interfaces.IFileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) GetAll() ([]*models.File, error) {
	return []*models.File{}, nil
}

func (r *fileRepository) GetByID(id uint) (*models.File, error) {
	return &models.File{}, nil
}

func (r *fileRepository) Create(file *models.File) error {
	return nil
}

func (r *fileRepository) Update(file *models.File) error {
	return nil
}

func (r *fileRepository) Delete(id uint) error {
	return nil
}
