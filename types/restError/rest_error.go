package resterror

import (
	"encoding/json"
	"net/http"
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

func NewMethodNotAllowedError() *RestError {
	return &RestError{
		Code: http.StatusMethodNotAllowed,
		Err:  "Method Not Allowed",
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "Bad Request",
	}
}

func NewBadRequestErrorWithCauses(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "Bad Request",
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "Internal Server Error",
	}
}

func NewInternalServerErrorWithCauses(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "Internal Server Error",
		Causes:  causes,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "Not Found",
	}
}

func NewNotFoundErrorWithCauses(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "Not Found",
		Causes:  causes,
	}
}
