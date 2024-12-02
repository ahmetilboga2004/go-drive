package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ahmetilboga2004/internal/application/dto"
	"github.com/ahmetilboga2004/internal/domain/interfaces"
	"github.com/ahmetilboga2004/internal/domain/models"
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

func (h *fileHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	files, err := h.fileService.GetAll()
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, "failed to get all files", err)
		return
	}

	filesDto := make([]dto.FileBasicInfo, len(files))
	for i, file := range files {
		filesDto[i] = file.ToBasicInfoDTO()
	}

	httphelper.SuccessResponse(w, http.StatusOK, "files retrieved successfully", filesDto)
}

func (h *fileHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "file id is required", nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid file id", err)
		return
	}

	file, err := h.fileService.GetByID(uint(id))

	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, "failed to get file", err)
		return
	}

	fileDto := file.ToDetailsDTO()

	httphelper.SuccessResponse(w, http.StatusOK, "file retrieved successfully", fileDto)
}

func (h *fileHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "file id is required", nil)
		return
	}

	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid file id", err)
		return
	}

	var fileDto dto.FileUpdate
	if err := json.NewDecoder(r.Body).Decode(&fileDto); err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid data", err)
		return
	}

	if err := httphelper.Validator.Struct(fileDto); err != nil {
		httphelper.ErrorResponse(w, http.StatusBadRequest, "invalid data", err)
		return
	}

	file := &models.File{
		Name:   fileDto.Name,
		Public: fileDto.Public,
	}

	err = h.fileService.Update(uint(idUint), file)
	if err != nil {
		httphelper.ErrorResponse(w, http.StatusInternalServerError, "failed to update file", err)
		return
	}
	httphelper.SuccessResponse(w, http.StatusOK, "file updated successfully", nil)

}
