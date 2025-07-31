package utils

import (
	"os"
	"testing"
)

func TestEncryptDecryptString(t *testing.T) {
	plaintext := "Hello, World!"
	passphrase := "securepassphrase"

	// Test encryption
	encrypted, err := EncryptString(plaintext, passphrase)
	if err != nil {
		t.Fatalf("Failed to encrypt string: %v", err)
	}

	// Test decryption
	decrypted, err := DecryptString(encrypted, passphrase)
	if err != nil {
		t.Fatalf("Failed to decrypt string: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decrypted text does not match original. Got: %s, Want: %s", decrypted, plaintext)
	}
}

func TestHashPassword(t *testing.T) {
	// Set up environment variable for PASSWORD_SALT
	passwordSalt := "random_salt_value"
	os.Setenv("PASSWORD_SALT", passwordSalt)
	defer os.Unsetenv("PASSWORD_SALT")

	password := "mypassword"

	// Test hashing
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Ensure hashed password is not empty
	if hashedPassword == "" {
		t.Error("Hashed password is empty")
	}

	// Decrypt the hashed password to verify
	decryptedPassword, err := DecryptString(hashedPassword, passwordSalt)
	if err != nil {
		t.Fatalf("Failed to decrypt hashed password: %v", err)
	}
	if decryptedPassword != password {
		t.Errorf("Decrypted password does not match original. Got: %s, Want: %s", decryptedPassword, password)
	}
}
