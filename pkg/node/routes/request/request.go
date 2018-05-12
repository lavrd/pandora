package request

import (
	"encoding/json"
	"io"

	"github.com/spacelavr/pandora/pkg/utils/http/response"
	"github.com/spacelavr/pandora/pkg/utils/validator"
)

// Candidate
type Candidate struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

// Validate validate incoming data for candidate
func (c *Candidate) Validate() *response.Error {
	switch {
	case c.Email == nil || !validator.IsEmail(*c.Email):
		return response.BadParameter("email")
	case c.Name == nil || len(*c.Name) == 0:
		return response.BadParameter("name")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for create account
func (c *Candidate) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(c); err != nil {
		return response.InvalidJSON()
	}
	return c.Validate()
}

// AccountFetch
type AccountFetch struct {
	PublicKey *string `json:"public_key"`
}

// Validate validate incoming data for fetch account
func (ac *AccountFetch) Validate() *response.Error {
	switch {
	case ac.PublicKey == nil || !validator.IsPublicKey(*ac.PublicKey):
		return response.BadParameter("public_key")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for fetch account
func (ac *AccountFetch) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(ac); err != nil {
		return response.InvalidJSON()
	}
	return ac.Validate()
}

// CertView
type CertView struct {
	Id *string `json:"id"`
}

// Validate validate incoming data for issue certificate
func (cv *CertView) Validate() *response.Error {
	switch {
	case cv.Id == nil || len(*cv.Id) == 0:
		return response.BadParameter("id")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for issue certificate
func (cv *CertView) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(cv); err != nil {
		return response.InvalidJSON()
	}
	return cv.Validate()
}

// CertificateIssue
type CertificateIssue struct {
	PublicKey   *string `json:"public_key"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate validate incoming data for issue certificate
func (ci *CertificateIssue) Validate() *response.Error {
	switch {
	case ci.PublicKey == nil || !validator.IsPublicKey(*ci.PublicKey):
		return response.BadParameter("public_key")
	case ci.Title == nil || len(*ci.Title) == 0:
		return response.BadParameter("title")
	case ci.Description == nil:
		*ci.Description = ""
		fallthrough
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate incoming data for issue certificate
func (ci *CertificateIssue) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(ci); err != nil {
		return response.InvalidJSON()
	}
	return ci.Validate()
}
