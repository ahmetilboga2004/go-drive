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
	"github.com/ahmetilboga2004/internal/infrastructure/utils/logger"
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
	files, err := s.fileRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *fileService) GetByID(id uint) (*models.File, error) {
	file, err := s.fileRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (s *fileService) Create(fileHeader *multipart.FileHeader, userId uint, isPublic bool) (*models.File, error) {
	if err := os.MkdirAll(s.uploadDir, 0755); err != nil {
		logger.Log.Sugar().Warnf("Error creating upload directory: %s", err)
		return nil, err
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), fileHeader.Filename)
	filePath := filepath.Join(s.uploadDir, fileName)

	filePathAbs, err := filepath.Abs(filePath)
	if err != nil {
		logger.Log.Sugar().Warnf("Error resolving absolute path: %s", err)
		return nil, err
	}

	logger.Log.Sugar().Infof("Uploading file to: %s", filePathAbs)

	src, err := fileHeader.Open()
	if err != nil {
		logger.Log.Sugar().Warnf("Error opening file: %s", err)
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(filePathAbs)
	if err != nil {
		logger.Log.Sugar().Warnf("Error creating file: %s", err)
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		logger.Log.Sugar().Warnf("Error copying file: %s", err)
		return nil, err
	}

	file := &models.File{
		Name:     fileHeader.Filename,
		Path:     filePathAbs,
		Size:     fileHeader.Size,
		FileType: filepath.Ext(fileHeader.Filename),
		Public:   isPublic,
		UserID:   userId,
	}

	if err := s.fileRepo.Create(file); err != nil {
		os.Remove(filePathAbs)
		logger.Log.Sugar().Warnf("Error saving file in DB: %s", err)
		return nil, err
	}

	return file, nil
}

func (s *fileService) Update(id uint, file *models.File) error {
	existingFile, err := s.fileRepo.GetByID(id)
	if err != nil {
		logger.Log.Sugar().Warnf("Error getting existing file: %s", err)
		return err
	}

	if file.Name != "" && file.Name != existingFile.Name {
		oldPath := existingFile.Path

		oldPathAbs, err := filepath.Abs(oldPath)
		if err != nil {
			logger.Log.Sugar().Warnf("Error resolving absolute path for old file: %s", err)
			return err
		}

		logger.Log.Sugar().Infof("Old file path: %s", oldPathAbs)

		if _, err := os.Stat(oldPathAbs); os.IsNotExist(err) {
			logger.Log.Sugar().Warnf("File not found: %s", oldPathAbs)
			return fmt.Errorf("file not found: %s", oldPathAbs)
		}

		ext := filepath.Ext(existingFile.Name)
		newFileName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), file.Name, ext)
		newPath := filepath.Join(s.uploadDir, newFileName)

		newPathAbs, err := filepath.Abs(newPath)
		if err != nil {
			logger.Log.Sugar().Warnf("Error resolving new path: %s", err)
			return err
		}

		logger.Log.Sugar().Infof("New file path: %s", newPathAbs)

		if err := os.Rename(oldPathAbs, newPathAbs); err != nil {
			logger.Log.Sugar().Warnf("Error renaming file: %s", err)
			return err
		}

		existingFile.Name = file.Name + ext
		existingFile.Path = newPathAbs
	}

	if file.Public != existingFile.Public {
		existingFile.Public = file.Public
	}

	if err := s.fileRepo.Update(existingFile); err != nil {
		logger.Log.Sugar().Warnf("Error updating file in DB: %s", err)
		return err
	}

	return nil
}

func (s *fileService) Delete(id uint) error {
	return nil
}
