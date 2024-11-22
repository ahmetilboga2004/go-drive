package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
)

type fileService struct {
	fileRepo  interfaces.IFileRepository
	uploadDir string
}

func NewFileService(fileRepo interfaces.IFileRepository) interfaces.IFileService {
	return &fileService{
		fileRepo:  fileRepo,
		uploadDir: "uploads",
	}
}

func (s *fileService) GetAll() ([]*models.File, error) {
	return []*models.File{}, nil
}

func (s *fileService) GetByID(id uint) (*models.File, error) {
	return &models.File{}, nil
}

func (s *fileService) Create(fileHeader *multipart.FileHeader, userId uint, isPublic bool) (*models.File, error) {
	if err := os.MkdirAll(s.uploadDir, 0755); err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	filePath := filepath.Join(s.uploadDir, fileName)

	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	file := &models.File{
		Name:     fileHeader.Filename,
		Path:     filePath,
		Size:     fileHeader.Size,
		FileType: filepath.Ext(fileHeader.Filename),
		Public:   isPublic,
		UserID:   userId,
	}

	if err := s.fileRepo.Create(file); err != nil {
		os.Remove(filePath)
		return nil, err
	}

	return file, nil
}

func (s *fileService) Update(id uint, file *models.File) error {
	return nil
}

func (s *fileService) Delete(id uint) error {
	return nil
}
