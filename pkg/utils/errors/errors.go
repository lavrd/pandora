package errors

import (
	"errors"
)

var (
	AlreadyExists = New("AlreadyExists")
	SendMailError = New("SendMailError")
	NotFound      = New("NotFound")
)

// New returns new errors with text
func New(text string) error {
	return errors.New(text)
}
