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
	Name  *string `json:"name"`
}

// Validate validate incoming data for sign up
func (s *SignUp) Validate() *errors.Response {
	switch {
	case s.Email == nil || !validator.IsEmail(*s.Email):
		return errors.BadParameter("email")
	case s.Type == nil || !validator.IsAccountType(*s.Type):
		return errors.BadParameter("type")
	case s.Name == nil || len(*s.Name) == 0:
		return errors.BadParameter("name")
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
		return errors.BadParameter("issuer_email")
	case ci.RecipientEmail == nil || !validator.IsEmail(*ci.RecipientEmail):
		return errors.BadParameter("recipient_email")
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
func (ci *CertificateIssue) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(ci); err != nil {
		return errors.InvalidJSON()
	}
	return ci.Validate()
}

// CertificateView
type CertificateView struct {
	Id *string `json:"id"`
}

// Validate validate incoming data for view certificate
func (cv *CertificateView) Validate() *errors.Response {
	switch {
	case cv.Id == nil || len(*cv.Id) == 0:
		return errors.BadParameter("id")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for view certificate
func (cv *CertificateView) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(cv); err != nil {
		return errors.InvalidJSON()
	}
	return cv.Validate()
}

// AccountVerify
type AccountVerify struct {
	Signature *string `json:"signature"`
	PublicKey *string `json:"public_key"`
}

// Validate validate incoming data for verify account
func (av *AccountVerify) Validate() *errors.Response {
	switch {
	case av.Signature == nil || !validator.IsSignature(*av.Signature):
		return errors.BadParameter("signature")
	case av.PublicKey == nil || !validator.IsPublicKey(*av.PublicKey):
		return errors.BadParameter("public_key")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for verify account
func (av *AccountVerify) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(av); err != nil {
		return errors.InvalidJSON()
	}
	return av.Validate()
}

// CertificateVerify
type CertificateVerify struct {
	Id                 *string `json:"id"`
	IssuerPublicKey    *string `json:"issuer_public_key"`
	RecipientPublicKey *string `json:"recipient_public_key"`
}

// Validate validate incoming data for verify certificate
func (cv *CertificateVerify) Validate() *errors.Response {
	switch {
	case cv.Id == nil || len(*cv.Id) == 0:
		return errors.BadParameter("id")
	case cv.IssuerPublicKey == nil || !validator.IsPublicKey(*cv.IssuerPublicKey):
		return errors.BadParameter("issuer_public_key")
	case cv.RecipientPublicKey == nil || !validator.IsPublicKey(*cv.RecipientPublicKey):
		return errors.BadParameter("recipient_public_key")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for verify certificate
func (cv *CertificateVerify) DecodeAndValidate(reader io.Reader) *errors.Response {
	if err := json.NewDecoder(reader).Decode(cv); err != nil {
		return errors.InvalidJSON()
	}
	return cv.Validate()
}
