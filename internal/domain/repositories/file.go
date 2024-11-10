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
	var files []*models.File
	if err := r.db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (r *fileRepository) GetByID(id uint) (*models.File, error) {
	var file models.File
	if err := r.db.Find(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *fileRepository) Create(file *models.File) error {
	result := r.db.Create(&file)
	return result.Error
}

func (r *fileRepository) Update(file *models.File) error {
	result := r.db.Model(&file).Updates(&file)
	return result.Error
}

func (r *fileRepository) Delete(id uint) error {
	result := r.db.Delete(&models.File{}, id)
	return result.Error
}
