package validator

import (
	"encoding/hex"
	"regexp"

	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/nacl"
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

// IsAccountType checking for valid account type
func IsAccountType(t int) bool {
	return t > 0 && t < types.MaxAccountType
}

// IsPassword checking for valid password
func IsPassword(password string) bool {
	return types.MinPasswordLen <= len(password) && len(password) <= types.MaxPasswordLen
}

// IsSignature checking for valid signature
func IsSignature(signature string) bool {
	return len(signature) == 512
}

func IsPublicKey(key string) bool {
	buf, _ := hex.DecodeString(key)

	switch {
	case nacl.PublicKeySize != len(buf):
		return false
	default:
		return true
	}
}
