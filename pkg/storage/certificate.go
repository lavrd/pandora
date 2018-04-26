package storage

import (
	"fmt"

	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
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

// CertificateFetch fetch certificate from storage
func (s *Storage) CertificateFetch(id string) (*types.Certificate, error) {
	var (
		cert  = &types.Certificate{}
		query = fmt.Sprintf(
			"for c in %s filter c.id == @id return c",
			CCertificate,
		)
		vars = map[string]interface{}{
			"id": id,
		}
	)

	_, err := s.Exec(query, vars, cert)
	if err != nil {
		if err == errors.DocumentNotFound {
			return nil, nil
		}
		return nil, err
	}

	return cert, nil
}
