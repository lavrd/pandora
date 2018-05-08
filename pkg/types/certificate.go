package types

import (
	"fmt"
	"time"
)

type (
	Certificate struct {
		Id        string           `json:"id"`
		Meta      *CertificateMeta `json:"meta"`
		Issuer    *Issuer          `json:"issuer"`
		Recipient *Recipient       `json:"recipient"`
	}

	CertificateMeta struct {
		Title       string    `json:"title"`
		Description string    `json:"description,omitempty"`
		Date        time.Time `json:"date"`
	}

	RecipientMeta struct {
		Name string `json:"name"`
	}

	IssuerMeta struct {
		Name string `json:"name"`
	}

	Recipient struct {
		Meta      *RecipientMeta `json:"meta"`
		Signature string         `json:"signature"`
	}
	Issuer struct {
		Meta      *IssuerMeta `json:"meta"`
		Signature string      `json:"signature"`
	}
)

// Public returns public certificate info
func (cert *Certificate) Public() *Certificate {
	return cert
}

func (cert *Certificate) Bytes() []byte {
	return []byte(fmt.Sprint(cert))
}
