package routes

import (
	"encoding/json"
	"io"

	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/validator"
)

type SignUp struct {
	Email *string `json:"email"`
}

func (s *SignUp) Validate() *errors.Response {
	switch {
	case !validator.IsEmail(*s.Email):
		return errors.BadParameter("email")
	default:
		return nil
	}
}

func (s *SignUp) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(s); err != nil {
		return errors.InvalidJSON()
	}
	return s.Validate()
}

type SignIn struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

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

func (s *SignIn) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(s); err != nil {
		return errors.InvalidJSON()
	}
	return s.Validate()
}

type AccountRecovery struct {
	Email *string
}

func (ar *AccountRecovery) Validate() *errors.Response {
	switch {
	case !validator.IsEmail(*ar.Email):
		return errors.BadParameter("email")
	default:
		return nil
	}
}

func (ar *AccountRecovery) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(ar); err != nil {
		return errors.InvalidJSON()
	}
	return ar.Validate()
}
