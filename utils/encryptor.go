package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keyLen   = 32      // AES-256
	saltLen  = 16      // tamanho do salt
	nonceLen = 12      // GCM standard nonce
	iter     = 100_000 // iterações PBKDF2 (ajustável para balancear segurança/performance)
)

// EncryptString recebe o texto plano, uma passphrase e opcionalmente um salt.
// Se salt == nil, ele gera um salt aleatório. Retorna base64(salt || nonce || ciphertext).
func EncryptString(plaintext, passphrase string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(passphrase), salt, iter, keyLen, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceLen)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aead.Seal(nil, nonce, []byte(plaintext), nil)

	// concat: salt || nonce || ciphertext
	full := append(append(salt, nonce...), ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(full)
	return encoded, nil
}

// DecryptString recebe o base64 produzido por EncryptString e a mesma passphrase.
// Retorna o plaintext ou erro.
func DecryptString(encoded, passphrase string) (string, error) {
	full, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	if len(full) < saltLen+nonceLen {
		return "", errors.New("dados criptografados muito curtos")
	}

	salt := full[:saltLen]
	nonce := full[saltLen : saltLen+nonceLen]
	ciphertext := full[saltLen+nonceLen:]

	key := pbkdf2.Key([]byte(passphrase), salt, iter, keyLen, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func HashPassword(password string) (string, error) {
	passwordSalt := os.Getenv("PASSWORD_SALT")
	if passwordSalt == "" {
		log.Fatal("PASSWORD_SALT is not set in the .env file")
	}

	hashedPassword, err := EncryptString(password, passwordSalt)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
