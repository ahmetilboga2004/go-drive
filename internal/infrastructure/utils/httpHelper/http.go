package httphelper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	trTranslations "github.com/go-playground/validator/v10/translations/tr"
)

var (
	Validator *validator.Validate
	trans     ut.Translator
)

func init() {
	Validator = validator.New()

	turkish := tr.New()
	uni := ut.New(turkish, turkish)
	trans, _ = uni.GetTranslator("tr")

	_ = trTranslations.RegisterDefaultTranslations(Validator, trans)
}

type apiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func ErrorResponse(w http.ResponseWriter, status int, message string, err error) {
	var errorDetails map[string]string

	if ve, ok := err.(validator.ValidationErrors); ok {
		errorDetails = make(map[string]string)
		for _, e := range ve {
			errorDetails[e.Field()] = e.Translate(trans)
		}
	} else if err != nil {
		errorDetails = map[string]string{"error": err.Error()}
	}

	response := apiResponse{
		Message: message,
		Error:   errorDetails,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
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
