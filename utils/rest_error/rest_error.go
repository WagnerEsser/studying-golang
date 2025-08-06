package rest_error

import (
	"encoding/json"
	"net/http"
	translations "studying-go/utils/translations"
)

type RestError struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (re *RestError) Error() string {
	return re.Message
}

func (re *RestError) Throw(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(re.Code)
	json.NewEncoder(w).Encode(re)
}

func NewRestError(message string, code int, err string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Code:    code,
		Err:     err,
		Causes:  causes,
	}
}

func NewMethodNotAllowedError(lang string) *RestError {
	return &RestError{
		Code: http.StatusMethodNotAllowed,
		Err:  translations.ErrorMessages["MethodNotAllowed"]["en-US"],
	}
}

func NewBadRequestError(message string, lang string) *RestError {
	if message == "" {
		message = translations.ErrorMessages["BadRequest"][lang]
	}
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     translations.ErrorMessages["BadRequest"]["en-US"],
	}
}

func NewBadRequestErrorWithCauses(message string, causes []Cause, lang string) *RestError {
	if message == "" {
		message = translations.ErrorMessages["BadRequest"][lang]
	}
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     translations.ErrorMessages["BadRequest"]["en-US"],
		Causes:  causes,
	}
}

func NewInternalServerError(message string, lang string) *RestError {
	if message == "" {
		message = translations.ErrorMessages["InternalServer"][lang]
	}
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     translations.ErrorMessages["InternalServer"]["en-US"],
	}
}

func NewInternalServerErrorWithCauses(message string, causes []Cause, lang string) *RestError {
	if message == "" {
		message = translations.ErrorMessages["InternalServer"][lang]
	}
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     translations.ErrorMessages["InternalServer"]["en-US"],
		Causes:  causes,
	}
}
func NewNotFoundError(message string, lang string) *RestError {
	if message == "" {
		message = translations.ErrorMessages["NotFound"][lang]
	}
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     translations.ErrorMessages["NotFound"]["en-US"],
	}
}

func NewNotFoundErrorWithCauses(message string, causes []Cause, lang string) *RestError {
	if message == "" {
		message = translations.ErrorMessages["NotFound"][lang]
	}
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     translations.ErrorMessages["NotFound"]["en-US"],
		Causes:  causes,
	}
}
