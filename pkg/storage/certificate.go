package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
)

const (
	CCertificate = "certificate"
)

// CertificateSave save certificate to storage
func (s *Storage) CertificateSave(cert *types.Certificate) error {
	_, err := s.Write(CCertificate, cert)
	if err != nil {
		return err
	}

	return nil
}
