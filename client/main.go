package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func main() {
	// Test /greet without query param
	resp, err := http.Get("http://localhost:8080/greet")

	if err != nil {
		fmt.Printf("Request failed: %v\n", err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Expected status code 200, got: %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err.Error())
		return
	}

	expected := "Hello, World!"
	if string(body) != expected {
		fmt.Printf("Expected response body '%s', got: '%s'\n", expected, string(body))
		return
	}

	fmt.Printf("Test /greet without query param passed. Response: %s\n", string(body))

	// Test /greet with query param
	resp, err = http.Get("http://localhost:8080/greet?name=Wagner")
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Expected status code 200, got: %d\n", resp.StatusCode)
		return
	}

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err.Error())
		return
	}

	expected = "Hello, Wagner!"
	if string(body) != expected {
		fmt.Printf("Expected response body '%s', got: '%s'\n", expected, string(body))
		return
	}

	fmt.Printf("Test /greet with query param passed. Response: %s\n", string(body))

	// Test /users
	resp, err = http.Get("http://localhost:8080/users")

	if err != nil {
		fmt.Printf("Request failed: %v\n", err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Expected status code 200, got: %d\n", resp.StatusCode)
		return
	}

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)

	if err != nil {
		fmt.Printf("Failed to unmarshal response body: %v\n", err.Error())
		return
	}

	if len(users) == 0 {
		fmt.Printf("Expected non-empty list of users, got empty list\n")
		return
	}

	fmt.Printf("Test /users passed. Response: %s\n", string(body))
}
