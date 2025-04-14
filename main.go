package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/greet", greetHandler)

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/new", createUser)
	http.HandleFunc("/users/", getUserByID)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const usersFile = "users.json"

func ensureUsersFileExists() {
	if _, err := os.Stat(usersFile); os.IsNotExist(err) {
		file, err := os.Create(usersFile)
		if err != nil {
			log.Fatalf("Failed to create users file: %v", err)
		}
		defer file.Close()

		file.Write([]byte("[]"))
	}
}

func readUsersFromFile() ([]User, error) {
	ensureUsersFileExists()

	data, err := os.ReadFile(usersFile)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func writeUsersToFile(users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(usersFile, data, 0644)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
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

	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
		return
	}

	users = append(users, newUser)

	err = writeUsersToFile(users)
	if err != nil {
		http.Error(w, "Failed to write to users file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateResponse{ID: newUser.ID.String()})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := readUsersFromFile()
	if err != nil {
		http.Error(w, "Failed to read users file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	users, err := readUsersFromFile()
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

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + name + "!"))
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Age         int       `json:"age"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     Address   `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	Number  int    `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
