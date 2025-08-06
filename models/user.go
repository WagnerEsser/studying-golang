package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	Age         int16     `json:"age" validate:"gte=0,lte=130"`          // example of age validation
	PhoneNumber string    `json:"phone_number" validate:"required,e164"` // e164 = international phone format
	Email       string    `json:"email" validate:"required,email,min=4,max=100"`
	Address     Address   `json:"address" validate:"required"`                  // dive = validate internal fields
	Password    string    `json:"password,omitempty" validate:"required,min=6"` // example
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Age         int16     `json:"age"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Address     Address   `json:"address"`
	Password    string    `json:"-"`
}

type Address struct {
	Street  string `json:"street" validate:"required"`
	Number  int    `json:"number" validate:"required,gt=0"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
