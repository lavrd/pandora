package types

import (
	"time"
)

// Certificate
type Certificate struct {
	Id        string           `json:"id"`
	Meta      *CertificateMeta `json:"meta"`
	Issuer    *Issuer          `json:"issuer"`
	Recipient *Recipient       `json:"recipient"`
}

// Public returns public certificate info
func (cert *Certificate) Public() *Certificate {
	return cert
}

// CertificateMeta
type CertificateMeta struct {
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date"`
}

// RecipientMeta
type RecipientMeta struct {
	Name string `json:"name"`
}

// IssuerMeta
type IssuerMeta struct {
	Name string `json:"name"`
}

// Recipient
type Recipient struct {
	Meta      *RecipientMeta `json:"meta"`
	Signature string         `json:"signature"`
}

// Issuer
type Issuer struct {
	Meta      *IssuerMeta `json:"meta"`
	Signature string      `json:"signature"`
}
