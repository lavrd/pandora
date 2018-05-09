package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

const (
	CCertificates = "CCertificates"
)

// CertificateSave save certificate to storage
func (s *Storage) CertificateSave(cert *types.Certificate) error {
	_, err := s.Write(CCertificates, cert)
	if err != nil {
		return err
	}

	return nil
}

// CertificateFetch fetch certificate from storage
func (s *Storage) CertificateFetch(id string) (*types.Certificate, error) {
	var (
		cert = &types.Certificate{}
	)

	_, err := s.Read(id, CCertificates, cert)
	if err != nil {
		if err == errors.NotFound {
			return nil, nil
		}
		return nil, err
	}

	return cert, nil
}
