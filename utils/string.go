package utils

import (
	"regexp"
)

// IsValidEmail checks if the input string is a valid email address.
//
// email string
// bool
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
