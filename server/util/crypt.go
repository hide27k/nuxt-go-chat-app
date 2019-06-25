package util

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates hased password.
// Returns string and error.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), errors.Wrap(err, "failed to generate from password")
}

// CheckHashOfPassword checks if given hash will be the same value of hashed password.
func CheckHashOfPassword(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
