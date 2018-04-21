package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
)

const (
	CollectionCertificate = "certificate"
)

// CertificateSave save certificate to storage
func (s *Storage) CertificateSave(cert *types.Certificate) error {
	_, err := s.Write(CollectionCertificate, cert)
	if err != nil {
		return err
	}

	return nil
}
