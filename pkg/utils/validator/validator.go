package validator

import (
	"regexp"

	"github.com/spacelavr/pandora/pkg/utils/log"
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

// IsPublicKey checking for valid public key
func IsPublicKey(key string) bool {
	return len(key) == 64
}
