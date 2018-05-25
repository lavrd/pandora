package arangodb

import (
	"github.com/spacelavr/pandora/pkg/pb"
)

const (
	CBlockchain = "CBlockchain"
)

func (s *ArangoDB) SaveCertBlock(block *pb.CertBlock) error {
	_, err := s.Write(CBlockchain, block)
	if err != nil {
		return err
	}

	return nil
}
