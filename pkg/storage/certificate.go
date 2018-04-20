package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
)

const (
	CollectionCertificate = "certificate"
)

// CertSave save certificate to storage
func (s *Storage) CertSave(cert *types.Certificate) error {
	_, err := s.Write(CollectionCertificate, cert)
	if err != nil {
		return err
	}

	return nil
}
