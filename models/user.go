package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	Age         int8      `json:"age"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     Address   `json:"address"`
	Password    string    `json:"password,omitempty"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Age         int8      `json:"age"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     Address   `json:"address"`
	Password    string    `json:"-"`
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
