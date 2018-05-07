package request

import (
	"encoding/json"
	"io"

	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/validator"
)

// Account
type Account struct {
	Email    *string `json:"email"`
	FullName *string `json:"full_name"`
}

// Validate validate incoming data for new account
func (a *Account) Validate() *errors.Response {
	switch {
	case a.Email == nil || !validator.IsEmail(*a.Email):
		return errors.BadParameter("email")
	case a.FullName == nil || len(*a.FullName) == 0:
		return errors.BadParameter("full_name")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for new account
func (a *Account) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(a); err != nil {
		return errors.InvalidJSON()
	}
	return a.Validate()
}
