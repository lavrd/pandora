package bcrypt

import (
	"github.com/spacelavr/pandora/pkg/utils/log"
	"golang.org/x/crypto/bcrypt"
)

// Encode encode password by bcrypt
func Encode(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return string(hash), nil
}

// Validate that hash and password are compared
func Validate(hashed, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			log.Error(err)
		}
		return err
	}

	return nil
}
