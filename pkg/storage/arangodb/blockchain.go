package arangodb

import (
	"github.com/spacelavr/pandora/pkg/pb"
)

const (
	CollectionBlockchain = "blockchain"
)

func (s *ArangoDB) SaveCertBlock(block *pb.CertBlock) error {
	_, err := s.Write(CollectionBlockchain, block)
	if err != nil {
		return err
	}

	return nil
}
