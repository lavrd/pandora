package distribution

import (
	"crypto/rsa"

	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/jwt"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/generator"
	"github.com/spacelavr/pandora/pkg/utils/mail"
)

func (d *Distribution) AccountCreate(opts *request.Account) (string, error) {
	acc, err := d.AccountFetchByEmail(*opts.Email)
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
	pri, pub := rsa.Encode(private, public)

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
	acc, err := d.AccountFetchByEmail(*opts.Email)
	if err != nil {
		return "", err
	}
	if acc == nil {
		return "", errors.DocumentNotFound
	}

	if err := bcrypt.Validate(acc.Secure.Password, *opts.Password); err != nil {
		return "", errors.InvalidCredentials
	}

	return jwt.New(acc)
}

// AccountRecovery recovery account, generate new password,
// save them and send recovery mail
func (d *Distribution) AccountRecovery(opts *request.AccountRecovery) error {
	acc, err := d.AccountFetchByEmail(*opts.Email)
	if err != nil {
		return err
	}
	if acc == nil {
		return errors.DocumentNotFound
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

// AccountVerify verify account by public key
func (d *Distribution) AccountVerify(opts *request.AccountVerify) (*types.Account, error) {
	pub, err := rsa.DecodePublic(*opts.PublicKey)
	if err != nil {
		return nil, err
	}

	if err := rsa.VerifyPSS(pub, *opts.Signature); err != nil {
		if err == rsa.ErrVerification {
			return nil, nil
		}
		return nil, err
	}

	acc, err := d.AccountFetchByPublic(*opts.PublicKey)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
