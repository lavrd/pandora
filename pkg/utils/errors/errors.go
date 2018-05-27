package errors

import (
	"errors"
)

var (
	AlreadyExists = New("AlreadyExists")
	SendMailError = New("SendMailError")
	NotFound      = New("NotFound")
)

// New returns new error
func New(text string) error {
	return errors.New(text)
}
