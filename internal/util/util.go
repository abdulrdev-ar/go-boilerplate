package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/inienam06/go-boilerplate/internal/exception"
)

// getGlobalSalt returns salt from environment variable
func getGlobalSalt() (string, error) {
	salt := os.Getenv("APP_PASSWORD_SALT")
	if salt == "" {
		return "", exception.NewInternalException("APP_PASSWORD_SALT is not set")
	}
	return salt, nil
}

// HashPassword hashes password using SHA-256 + global salt from ENV
func HashPassword(password string) (string, error) {
	salt, err := getGlobalSalt()
	if err != nil {
		return "", err
	}

	combined := password + salt
	hash := sha256.Sum256([]byte(combined))

	return hex.EncodeToString(hash[:]), nil
}

// VerifyPassword compares raw password with stored hash
func VerifyPassword(rawPassword, storedHash string) (bool, error) {
	hash, err := HashPassword(rawPassword)
	if err != nil {
		return false, err
	}

	return hash == storedHash, nil
}
