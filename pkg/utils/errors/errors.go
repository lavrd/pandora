package errors

import (
	"errors"
)

var (
	AlreadyExists = New("already exists")
	SendMailError = New("send mail error")
	NotFound      = New("not found")
)

// New returns new error
func New(text string) error {
	return errors.New(text)
}
