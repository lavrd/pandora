package request

import (
	"encoding/json"
	"io"

	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/validator"
)

// AccountCreate
type AccountCreate struct {
	Email    *string `json:"email"`
	FullName *string `json:"full_name"`
}

// Validate validate incoming data for create account
func (ac *AccountCreate) Validate() *errors.Response {
	switch {
	case ac.Email == nil || !validator.IsEmail(*ac.Email):
		return errors.BadParameter("email")
	case ac.FullName == nil || len(*ac.FullName) == 0:
		return errors.BadParameter("full_name")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for create account
func (ac *AccountCreate) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(ac); err != nil {
		return errors.InvalidJSON()
	}
	return ac.Validate()
}

// AccountFetch
type AccountFetch struct {
	PublicKey *string `json:"public_key"`
}

// Validate validate incoming data for fetch account
func (ac *AccountFetch) Validate() *errors.Response {
	switch {
	case ac.PublicKey == nil || !validator.IsPublicKey(*ac.PublicKey):
		return errors.BadParameter("public_key")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for fetch account
func (ac *AccountFetch) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(ac); err != nil {
		return errors.InvalidJSON()
	}
	return ac.Validate()
}
