package main

import (
	"log"
	"net/http"
	"studying-go/handlers"
)

func main() {
	http.HandleFunc("/", handlers.GreetHandler)
	http.HandleFunc("/greet", handlers.GreetHandler)

	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/", handlers.GetUserByID)
	http.HandleFunc("/users/new", handlers.CreateUser)
	http.HandleFunc("/users/delete/", handlers.DeleteUser)
	http.HandleFunc("/users/edit/", handlers.UpdateUser)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
