package validator

import (
	"regexp"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/types"
)

// IsEmail checking for valid email
func IsEmail(email string) bool {
	const pattern = "[^@]+@[^\\.]+\\..+"
	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		log.Error(err)
	}
	return matched
}

// IsAccountType checking for valid account type
func IsAccountType(t int) bool {
	return t > 0 && t < types.MaxAccountType
}

// IsPassword checking for valid password
func IsPassword(password string) bool {
	return types.MinPasswordLen <= len(password) && len(password) <= types.MaxPasswordLen
}
