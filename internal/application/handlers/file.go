package handlers

import (
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/go-playground/validator/v10"
)

type fileHandler struct {
	fileService interfaces.IFileService
	validator   *validator.Validate
}

func NewFileHandler(fileService interfaces.IFileService) *fileHandler {
	return &fileHandler{
		fileService: fileService,
		validator:   validator.New(),
	}
}
