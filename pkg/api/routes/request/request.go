package request

import (
	"encoding/json"
	"io"

	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/validator"
)

// SignUp
type SignUp struct {
	Email *string `json:"email"`
	Type  *int    `json:"type"`
}

// Validate validate incoming data for sign up
func (s *SignUp) Validate() *errors.Response {
	switch {
	case !validator.IsEmail(*s.Email):
		return errors.BadParameter("email")
	case !validator.IsAccountType(*s.Type):
		return errors.BadParameter("type")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for sign up
func (s *SignUp) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(s); err != nil {
		return errors.InvalidJSON()
	}
	return s.Validate()
}

// SignIn
type SignIn struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// Validate validate incoming data for sign in
func (s *SignIn) Validate() *errors.Response {
	switch {
	case !validator.IsEmail(*s.Email):
		return errors.BadParameter("email")
	case !validator.IsPassword(*s.Password):
		return errors.BadParameter("password")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for sign in
func (s *SignIn) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(s); err != nil {
		return errors.InvalidJSON()
	}
	return s.Validate()
}

// AccountRecovery
type AccountRecovery struct {
	Email *string
}

// Validate validate incoming data for account recovery
func (ar *AccountRecovery) Validate() *errors.Response {
	switch {
	case !validator.IsEmail(*ar.Email):
		return errors.BadParameter("email")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for account recovery
func (ar *AccountRecovery) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(ar); err != nil {
		return errors.InvalidJSON()
	}
	return ar.Validate()
}
