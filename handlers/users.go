package handlers

import (
	"encoding/json"
	"net/http"
	"studying-go/models"
	"studying-go/storage"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)
		return
	}
	newUser.ID = id

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
		return
	}

	users = append(users, newUser)

	err = storage.WriteUsersToFile(users)
	if err != nil {
		http.Error(w, "Failed to write to users file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateResponse{ID: newUser.ID.String()})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/users/edit/"):]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
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
		http.Error(w, "Failed to write to users file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/users/delete/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
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
		http.Error(w, "Failed to write to users file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	users, err := storage.ReadUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
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

	http.Error(w, "User not found", http.StatusNotFound)
}
