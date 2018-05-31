package errors

import (
	"errors"
)

var (
	ErrAlreadyExists = New("error already exists")
	ErrMail          = New("mail error")
	ErrNotFound      = New("error not found")
)

// New returns new error
func New(text string) error {
	return errors.New(text)
}
