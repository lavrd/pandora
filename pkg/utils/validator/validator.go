package validator

import (
	"regexp"
)

// IsEmail checking for valid email
func IsEmail(email string) bool {
	const pattern = "[^@]+@[^\\.]+\\..+"
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// IsPublicKey checking for valid public key
func IsPublicKey(key string) bool {
	return len(key) == 64
}
