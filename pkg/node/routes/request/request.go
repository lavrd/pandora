package request

import (
	"encoding/json"
	"io"

	"pandora/pkg/utils/http/response"
	"pandora/pkg/utils/validator"
)

// Candidate
type Candidate struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

// Validate validate candidate request
func (c *Candidate) Validate() *response.Response {
	switch {
	case c.Email == nil || !validator.IsEmail(*c.Email):
		return response.BadParameter("email")
	case c.Name == nil || len(*c.Name) == 0:
		return response.BadParameter("name")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate candidate request
func (c *Candidate) DecodeAndValidate(reader io.Reader) *response.Response {
	if err := json.NewDecoder(reader).Decode(c); err != nil {
		return response.InvalidJSON("email, name")
	}
	return c.Validate()
}

// MemberFetch
type MemberFetch struct {
	PublicKey *string `json:"public_key"`
}

// Validate validate member fetch request
func (mf *MemberFetch) Validate() *response.Response {
	switch {
	case mf.PublicKey == nil || !validator.IsPublicKey(*mf.PublicKey):
		return response.BadParameter("public_key")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate member fetch request
func (mf *MemberFetch) DecodeAndValidate(reader io.Reader) *response.Response {
	if err := json.NewDecoder(reader).Decode(mf); err != nil {
		return response.InvalidJSON("public_key")
	}
	return mf.Validate()
}

// CertView
type CertView struct {
	Id *string `json:"id"`
}

// Validate validate cert view request
func (cv *CertView) Validate() *response.Response {
	switch {
	case cv.Id == nil || len(*cv.Id) == 0:
		return response.BadParameter("id")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate cert view request
func (cv *CertView) DecodeAndValidate(reader io.Reader) *response.Response {
	if err := json.NewDecoder(reader).Decode(cv); err != nil {
		return response.InvalidJSON("id")
	}
	return cv.Validate()
}

// CertIssue
type CertIssue struct {
	PublicKey   *string `json:"public_key"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate validate cert issue request
func (ci *CertIssue) Validate() *response.Response {
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

// DecodeAndValidate decode and validate cert issue request
func (ci *CertIssue) DecodeAndValidate(reader io.Reader) *response.Response {
	if err := json.NewDecoder(reader).Decode(ci); err != nil {
		return response.InvalidJSON("public_key, title, description")
	}
	return ci.Validate()
}

// CertVerify
type CertVerify struct {
	Id *string `json:"id"`
}

// Validate validate cert verify request
func (cv *CertVerify) Validate() *response.Response {
	switch {
	case cv.Id == nil || len(*cv.Id) == 0:
		return response.BadParameter("id")
	default:
		return nil
	}
}

// DecodeAndValidate decode and validate cert verify request
func (cv *CertVerify) DecodeAndValidate(reader io.Reader) *response.Response {
	if err := json.NewDecoder(reader).Decode(cv); err != nil {
		return response.InvalidJSON("id")
	}
	return cv.Validate()
}
