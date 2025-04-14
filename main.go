package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/greet", greetHandler)

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/new", createUser)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var users []User

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

	id, err = uuid.NewRandom()
	if err != nil {
		http.Error(w, "Failed to generate UUID", http.StatusInternalServerError)
		return
	}
	newUser.Address.ID = id

	users = append(users, newUser)
	log.Printf("Current users: %+v\n", users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User created successfully"}`))
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	usersMock := []User{
		{
			Age:         30,
			PhoneNumber: "(21) 99876-5432",
			Email:       "bob@exemplo.com",
			Address: Address{
				Street:  "Avenida Atlântica",
				Number:  456,
				City:    "Rio de Janeiro",
				State:   "RJ",
				Country: "Brasil",
			},
		},
		{
			Age:         35,
			PhoneNumber: "(31) 95555-5555",
			Email:       "charlie@exemplo.com",
			Address: Address{
				Street:  "Praça da Liberdade",
				Number:  789,
				City:    "Belo Horizonte",
				State:   "MG",
				Country: "Brasil",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersMock)
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Age         int       `json:"age"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     Address   `json:"address"`
}

type Address struct {
	ID      uuid.UUID `json:"id"`
	Street  string    `json:"street"`
	Number  int       `json:"number"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	Country string    `json:"country"`
}
