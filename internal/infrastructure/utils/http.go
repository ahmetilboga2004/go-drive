package utils

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, status int, message string, data any) {
	response := apiResponse{
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, status int, message string, err error) {
	response := apiResponse{
		Message: message,
		Error:   err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
