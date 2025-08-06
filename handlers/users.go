package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"studying-go/models"
	"studying-go/storage"
	"studying-go/utils/encryptor"
	restError "studying-go/utils/rest_error"
	"studying-go/utils/translations"
	"studying-go/utils/validator"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en-US"
	}

	if r.Method != http.MethodPost {
		restError.NewMethodNotAllowedError(lang).Throw(w)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		slog.Error("Failed to decode request body", "error", err)
		restError.NewBadRequestError("", lang).Throw(w)
		return
	}

	errCauses := validator.ValidateStruct(newUser, lang)
	if errCauses != nil {
		restError.NewBadRequestErrorWithCauses("", errCauses, lang).Throw(w)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		slog.Error("Failed to generate UUID", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToGenerateUUID"][lang], lang).Throw(w)
		return
	}

	newUser.ID = id
	encryptedPassword, err := encryptor.HashPassword(newUser.Password)
	if err != nil {
		slog.Error("Failed to encrypt password", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToEncryptPassword"][lang], lang).Throw(w)
		return
	}
	newUser.Password = encryptedPassword

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToReadUsersFile"][lang], lang).Throw(w)
		return
	}

	users = append(users, newUser)

	err = storage.WriteUsersToFile(users)
	if err != nil {
		slog.Error("Failed to write users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToWriteUsersFile"][lang], lang).Throw(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateResponse{ID: newUser.ID.String()})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en-US"
	}

	if r.Method != http.MethodPut {
		restError.NewMethodNotAllowedError(lang).Throw(w)
		return
	}

	var updatedUser models.UserResponse
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		slog.Error("Failed to decode request body", "error", err)
		restError.NewBadRequestError("", lang).Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToReadUsersFile"][lang], lang).Throw(w)
		return
	}

	for i, user := range users {
		if user.ID == updatedUser.ID {
			users[i].Email = updatedUser.Email
			users[i].Age = updatedUser.Age
			users[i].PhoneNumber = updatedUser.PhoneNumber
			users[i].Address = updatedUser.Address

			errCauses := validator.ValidateStruct(users[i], lang)
			if errCauses != nil {
				restError.NewBadRequestErrorWithCauses("", errCauses, lang).Throw(w)
				return
			}
			break
		}
	}

	err = storage.WriteUsersToFile(users)
	if err != nil {
		slog.Error("Failed to write users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToWriteUsersFile"][lang], lang).Throw(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en-US"
	}

	if r.Method != http.MethodDelete {
		restError.NewMethodNotAllowedError(lang).Throw(w)
		return
	}

	idStr := r.URL.Path[len("/users/delete/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.Error("Failed to parse user ID", "error", err)
		restError.NewBadRequestError(translations.ErrorMessages["InvalidUserID"][lang], lang).Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToReadUsersFile"][lang], lang).Throw(w)
		return
	}

	var updatedUsers []models.User
	for _, user := range users {
		if user.ID != id {
			updatedUsers = append(updatedUsers, user)
		}
	}

	err = storage.WriteUsersToFile(updatedUsers)
	if err != nil {
		slog.Error("Failed to write users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToWriteUsersFile"][lang], lang).Throw(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en-US"
	}

	if r.Method != http.MethodGet {
		restError.NewMethodNotAllowedError(lang).Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFileToResponse()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToReadUsersFile"][lang], lang).Throw(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		lang = "en-US"
	}

	if r.Method != http.MethodGet {
		restError.NewMethodNotAllowedError(lang).Throw(w)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.Error("Failed to parse user ID", "error", err)
		restError.NewBadRequestError(translations.ErrorMessages["InvalidUserID"][lang], lang).Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFileToResponse()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError(translations.ErrorMessages["FailedToReadUsersFile"][lang], lang).Throw(w)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	slog.Info("User not found", "userID", id)
	restError.NewNotFoundError(translations.ErrorMessages["UserNotFound"][lang], lang).Throw(w)
}
