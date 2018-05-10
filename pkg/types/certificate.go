package types

import (
	"encoding/hex"
	"fmt"
	"time"
)

type (
	Certificate struct {
		Key       string           `json:"_key"`
		Id        string           `json:"id"`
		Meta      *CertificateMeta `json:"meta"`
		Issuer    *Issuer          `json:"issuer"`
		Recipient *Recipient       `json:"recipient"`
	}

	CertificateMeta struct {
		Title       string    `json:"title"`
		Description string    `json:"description,omitempty"`
		Timestamp   time.Time `json:"timestamp"`
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
	buf, _ := hex.DecodeString(fmt.Sprint(cert))
	return buf
}
