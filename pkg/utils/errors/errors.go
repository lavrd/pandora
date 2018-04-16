package errors

import (
	"errors"
)

var (
	AccountAlreadyExists = New("account already exists")
	AccountNotFound      = New("account not found")
	InvalidCredentials   = New("invalid credentials")
	SendMailError        = New("send mail error")
	PemBlockParseFailed  = New("failed to parse pem block")
)

// New returns new errors with text
func New(text string) error {
	return errors.New(text)
}
