package validator

import (
	"regexp"

	"github.com/spacelavr/pandora/pkg/types"
)

// IsEmail checking for valid email
func IsEmail(email string) bool {
	const pattern = "[^@]+@[^\\.]+\\..+"
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// IsPassword checking for valid password
func IsPassword(password string) bool {
	return types.MinPasswordLen <= len(password) && len(password) <= types.MaxPasswordLen
}
