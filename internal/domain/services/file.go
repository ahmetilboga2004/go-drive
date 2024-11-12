package services

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
)

type fileService struct {
	fileRepo interfaces.IFileRepository
}

func NewFileService(fileRepo interfaces.IFileRepository) interfaces.IFileService {
	return &fileService{fileRepo: fileRepo}
}

func (s *fileService) GetAll() ([]*models.File, error) {
	return []*models.File{}, nil
}

func (s *fileService) GetByID(id uint) (*models.File, error) {
	return &models.File{}, nil
}

func (s *fileService) Create(file *models.File) error {
	return nil
}

func (s *fileService) Update(id uint, file *models.File) error {
	return nil
}

func (s *fileService) Delete(id uint) error {
	return nil
}
