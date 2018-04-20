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
	case s.Email == nil || !validator.IsEmail(*s.Email):
		return errors.BadParameter("email")
	case s.Type == nil || !validator.IsAccountType(*s.Type):
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
	case s.Email == nil || !validator.IsEmail(*s.Email):
		return errors.BadParameter("email")
	case s.Password == nil || !validator.IsPassword(*s.Password):
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
	Email *string `json:"email"`
}

// Validate validate incoming data for account recovery
func (ar *AccountRecovery) Validate() *errors.Response {
	switch {
	case ar.Email == nil || !validator.IsEmail(*ar.Email):
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

// CertificateIssue
type CertificateIssue struct {
	IssuerEmail    *string `json:"issuer_email"`
	RecipientEmail *string `json:"recipient_email"`
	Title          *string `json:"title"`
	Description    *string `json:"description"`
}

// Validate validate incoming data for issue certificate
func (ci *CertificateIssue) Validate() *errors.Response {
	switch {
	case ci.IssuerEmail == nil || !validator.IsEmail(*ci.IssuerEmail):
		return errors.BadParameter("issuer email")
	case ci.RecipientEmail == nil || !validator.IsEmail(*ci.RecipientEmail):
		return errors.BadParameter("recipient email")
	case ci.Title == nil || len(*ci.Title) == 0:
		return errors.BadParameter("title")
	case ci.Description == nil:
		*ci.Description = ""
		fallthrough
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for issue certificate
func (si *CertificateIssue) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(si); err != nil {
		return errors.InvalidJSON()
	}
	return si.Validate()
}
