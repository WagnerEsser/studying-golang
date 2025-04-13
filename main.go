package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", greetHandler)
	http.HandleFunc("/greet", greetHandler)

	http.HandleFunc("/users", getUsers)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
			Age:         25,
			PhoneNumber: "(11) 91234-5678",
			Email:       "alice@exemplo.com",
			Address: Address{
				Street:  "Rua das Flores",
				Number:  123,
				City:    "São Paulo",
				State:   "SP",
				Country: "Brasil",
			},
		},
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
	ID          int     `json:"id"`
	Age         int     `json:"age"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Address     Address `json:"address"`
}

type Address struct {
	ID      int    `json:"id"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}
