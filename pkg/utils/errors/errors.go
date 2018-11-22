package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrAlreadyExists = New("error already exists")
	ErrMail          = New("mail error")
	ErrNotFound      = New("error not found")
)

// WithStack return new error with stack trace
func WithStack(err error) error {
	return errors.WithStack(err)
}

// New creates new error with stack trace
func New(message string) error {
	return errors.WithStack(errors.New(message))
}
