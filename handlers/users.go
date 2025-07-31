package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"studying-go/models"
	"studying-go/storage"
	restError "studying-go/types"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		restError.NewMethodNotAllowedError().Throw(w)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		slog.Error("Failed to decode request body", "error", err)
		restError.NewBadRequestError("Invalid request body").Throw(w)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		slog.Error("Failed to generate UUID", "error", err)
		restError.NewInternalServerError("Failed to generate UUID").Throw(w)
		return
	}
	newUser.ID = id

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError("Failed to read users file").Throw(w)
		return
	}

	users = append(users, newUser)

	err = storage.WriteUsersToFile(users)
	if err != nil {
		slog.Error("Failed to write users file", "error", err)
		restError.NewInternalServerError("Failed to write to users file").Throw(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateResponse{ID: newUser.ID.String()})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		restError.NewMethodNotAllowedError().Throw(w)
		return
	}

	idStr := r.URL.Path[len("/users/edit/"):]

	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.Error("Failed to parse user ID", "error", err)
		restError.NewBadRequestError("Invalid user ID").Throw(w)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		slog.Error("Failed to decode request body", "error", err)
		restError.NewBadRequestError("Invalid request body").Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError("Failed to read users file").Throw(w)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			break
		}
	}

	err = storage.WriteUsersToFile(users)
	if err != nil {
		slog.Error("Failed to write users file", "error", err)
		restError.NewInternalServerError("Failed to write to users file").Throw(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		restError.NewMethodNotAllowedError().Throw(w)
		return
	}

	idStr := r.URL.Path[len("/users/delete/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.Error("Failed to parse user ID", "error", err)
		restError.NewBadRequestError("Invalid user ID").Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError("Failed to read users file").Throw(w)
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
		restError.NewInternalServerError("Failed to write users file").Throw(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		restError.NewMethodNotAllowedError().Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError("Failed to read users file").Throw(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		restError.NewMethodNotAllowedError().Throw(w)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.Error("Failed to parse user ID", "error", err)
		restError.NewBadRequestError("Invalid user ID").Throw(w)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		slog.Error("Failed to read users file", "error", err)
		restError.NewInternalServerError("Failed to read users file").Throw(w)
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
	restError.NewNotFoundError("User not found").Throw(w)
}
