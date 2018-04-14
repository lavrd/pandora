package distribution

import (
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/bcrypt"
	"github.com/spacelavr/pandora/pkg/utils/crypto/jwt"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/generator"
	"github.com/spacelavr/pandora/pkg/utils/mail"
)

// CreateAccount create account
func CreateAccount(email string) (string, error) {
	acc, err := storage.AccountFetch(email)
	if err != nil {
		return "", err
	}
	if acc != nil {
		return "", errors.AccountAlreadyExists
	}

	password := generator.Password()
	hashed, err := bcrypt.Encode(password)
	if err != nil {
		return "", err
	}

	acc = &types.Account{
		Email:    email,
		Password: hashed,
	}

	if err = storage.AccountPut(acc); err != nil {
		return "", err
	}

	if err = mail.SendAccountCreated(email, password); err != nil {
		return "", err
	}

	return jwt.New(acc);
}

func CreateSession(email, password string) (string, error) {
	acc, err := storage.AccountFetch(email)
	if err != nil {
		return "", err
	}
	if acc == nil {
		return "", errors.AccountNotFound
	}

	if err := bcrypt.Validate(acc.Password, password); err != nil {
		return "", errors.InvalidCredentials
	}

	return jwt.New(acc)
}

func AccountRecovery(email string) error {
	acc, err := storage.AccountFetch(email)
	if err != nil {
		return err
	}
	if acc == nil {
		return errors.AccountNotFound
	}

	password := generator.Password()
	hashed, err := bcrypt.Encode(password)
	if err != nil {
		return err
	}

	acc.Password = hashed

	err = storage.AccountPut(acc)
	if err != nil {
		return err
	}

	return mail.SendAccountRecovery(email, password)
}
