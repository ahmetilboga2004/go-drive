package services

import "github.com/ahmetilboga2004/internal/domain/interfaces"

type fileService struct {
	fileRepo interfaces.IFileRepository
}

func NewFileService(fileRepo interfaces.IFileRepository) *fileService {
	return &fileService{fileRepo: fileRepo}
}
