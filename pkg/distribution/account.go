package distribution

import (
	"github.com/spacelavr/pandora/pkg/api/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/bcrypt"
	"github.com/spacelavr/pandora/pkg/utils/crypto/jwt"
	"github.com/spacelavr/pandora/pkg/utils/crypto/rsa"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/generator"
	"github.com/spacelavr/pandora/pkg/utils/mail"
)

// AccountCreate generate password, keys, create account, save them,
// send mail with credentials and returns jwt token
func (d *Distribution) AccountCreate(opts *request.SignUp) (string, error) {
	acc, err := d.AccountFetch(*opts.Email)
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

	private, public, err := rsa.GenerateKeys()
	if err != nil {
		return "", err
	}
	pri, pub := rsa.Marshal(private, public)

	acc = &types.Account{
		Meta: &types.AccountMeta{
			Email: *opts.Email,
			Type:  *opts.Type,
			Name:  *opts.Name,
		},
		PublicKey: pub,
		Secure: &types.AccountSecure{
			Password:   hashed,
			PrivateKey: pri,
		},
	}

	if err = d.AccountSave(acc); err != nil {
		return "", err
	}

	if err = mail.SendAccountCreated(*opts.Email, password); err != nil {
		return "", err
	}

	return jwt.New(acc)
}

// SessionNew create new session and returns jwt token
func (d *Distribution) SessionNew(opts *request.SignIn) (string, error) {
	acc, err := d.AccountFetch(*opts.Email)
	if err != nil {
		return "", err
	}
	if acc == nil {
		return "", errors.AccountNotFound
	}

	if err := bcrypt.Validate(acc.Secure.Password, *opts.Password); err != nil {
		return "", errors.InvalidCredentials
	}

	return jwt.New(acc)
}

// AccountRecovery recovery account, generate new password,
// save them and send recovery mail
func (d *Distribution) AccountRecovery(opts *request.AccountRecovery) error {
	acc, err := d.AccountFetch(*opts.Email)
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

	acc.Secure.Password = hashed

	err = d.AccountUpdate(acc)
	if err != nil {
		return err
	}

	return mail.SendAccountRecovery(*opts.Email, password)
}
