package storage

import (
	"encoding/json"
	"log"
	"os"
	"studying-go/models"
)

const usersFile = "storage/users.json"

func EnsureUsersFileExists() {
	if _, err := os.Stat(usersFile); os.IsNotExist(err) {
		file, err := os.Create(usersFile)
		if err != nil {
			log.Fatalf("Failed to create users file: %v", err)
		}
		defer file.Close()

		file.Write([]byte("[]"))
	}
}

func ReadUsersFromFile() ([]models.User, error) {
	EnsureUsersFileExists()

	data, err := os.ReadFile(usersFile)
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func WriteUsersToFile(users []models.User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(usersFile, data, 0644)
}
