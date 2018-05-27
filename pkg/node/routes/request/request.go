package request

import (
	"encoding/json"
	"io"

	"github.com/spacelavr/pandora/pkg/utils/http/response"
	"github.com/spacelavr/pandora/pkg/utils/validator"
)

type Candidate struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

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

func (c *Candidate) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(c); err != nil {
		return response.InvalidJSON()
	}
	return c.Validate()
}

type MemberFetch struct {
	PublicKey *string `json:"public_key"`
}

func (mf *MemberFetch) Validate() *response.Error {
	switch {
	case mf.PublicKey == nil || !validator.IsPublicKey(*mf.PublicKey):
		return response.BadParameter("public_key")
	default:
		return nil
	}
}

func (mf *MemberFetch) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(mf); err != nil {
		return response.InvalidJSON()
	}
	return mf.Validate()
}

type CertView struct {
	Id *string `json:"id"`
}

func (cv *CertView) Validate() *response.Error {
	switch {
	case cv.Id == nil || len(*cv.Id) == 0:
		return response.BadParameter("id")
	default:
		return nil
	}
}

func (cv *CertView) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(cv); err != nil {
		return response.InvalidJSON()
	}
	return cv.Validate()
}

type CertificateIssue struct {
	PublicKey   *string `json:"public_key"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

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

func (ci *CertificateIssue) DecodeAndValidate(reader io.Reader) *response.Error {
	if err := json.NewDecoder(reader).Decode(ci); err != nil {
		return response.InvalidJSON()
	}
	return ci.Validate()
}
