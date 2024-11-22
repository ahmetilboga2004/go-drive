package handlers

import (
	"net/http"

	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/infrastructure/utils/auth"
	httphelper "github.com/ahmetilboga2004/internal/infrastructure/utils/httpHelper"
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

func (h *fileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "Failed to parse form", err)
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "Failed to get file from form", err)
		return
	}
	defer file.Close()

	isPublic := r.FormValue("public") == "true"

	userID, err := auth.GetUserIDFromContext(r)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusUnauthorized, "unauthorized", err)
		return
	}
	uploadedFile, err := h.fileService.Create(fileHeader, userID, isPublic)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, "failed to upload file", err)
		return
	}

	uploadedFileData := uploadedFile.ToBasicInfoDTO()

	httphelper.SuccessResponse(w, http.StatusOK, "file uploaded successfully", uploadedFileData)
}
