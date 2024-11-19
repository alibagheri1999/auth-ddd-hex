package utils

import (
	"net/mail"
	"unicode"
)

// IsValidEmail validates the email format using regex
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidPassword checks if a password meets the requirements:
// - At least one lowercase letter
// - At least one uppercase letter
// - At least one digit
// - At least one special character
// - Minimum length of 8 characters
func IsValidPassword(password string) bool {
	var hasUpper, hasLower, hasDigit, hasSpecial bool

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
