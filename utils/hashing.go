package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// function for hassing password
func HashPassword(password string) (string, error) {
	// Hash the password using bcrypt with a cost factor of 12
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
