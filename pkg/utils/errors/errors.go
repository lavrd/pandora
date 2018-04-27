package errors

import (
	"errors"
)

var (
	AccountAlreadyExists    = New("AccountAlreadyExists")
	InvalidCredentials      = New("InvalidCredentials")
	SendMailError           = New("SendMailError")
	PemBlockParseFailed     = New("PemBlockParseFailed")
	UnexpectedSigningMethod = New("UnexpectedSigningMethod")
	DocumentNotFound        = New("DocumentNotFound")
	IssueCertToNonRecipient = New("IssueCertToNonRecipient")
)

// New returns new errors with text
func New(text string) error {
	return errors.New(text)
}
