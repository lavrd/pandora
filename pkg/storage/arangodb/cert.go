package arangodb

import (
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

const (
	CollectionCertificate = "certificate"
)

// CertificateSave save certificate
func (s *ArangoDB) CertificateSave(cert *pb.Cert) error {
	_, err := s.Write(CollectionCertificate, cert)
	if err != nil {
		return err
	}

	return nil
}

// CertificateFetch fetch certificate
func (s *ArangoDB) CertificateFetch(id string) (*pb.Cert, error) {
	var (
		cert = &pb.Cert{}
	)

	_, err := s.Read(id, CollectionCertificate, cert)
	if err != nil {
		if err == errors.NotFound {
			return nil, nil
		}
		return nil, err
	}

	return cert, nil
}
