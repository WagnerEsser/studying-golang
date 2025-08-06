package routes

import (
	"net/http"
	"studying-go/handlers"
)

func InitRoutes() {
	http.HandleFunc("/", handlers.GreetHandler)
	http.HandleFunc("/greet", handlers.GreetHandler)

	// users
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/users/", handlers.GetUserByID)
	http.HandleFunc("/users/new", handlers.CreateUser)
	http.HandleFunc("/users/delete/", handlers.DeleteUser)
	http.HandleFunc("/users/edit", handlers.UpdateUser)
}
