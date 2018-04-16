package types

import (
	"crypto/rsa"
	"time"
)

// Certificate
type Certificate struct {
	Meta      *Meta      `json:"meta"`
	Issuer    *Issuer    `json:"issuer"`
	Recipient *Recipient `json:"recipient"`
}

// Meta
type Meta struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

// Recipient
type Recipient struct {
	FullName  string         `json:"full_name"`
	PublicKey *rsa.PublicKey `json:"public_key"`
}

// Issuer
type Issuer struct {
	PublicKey *rsa.PublicKey `json:"public_key"`
	Name      string         `json:"name"`
}

// Public returns public certificate info
func (cert *Certificate) Public() *Certificate {
	return cert
}
