package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
)

const (
	CBlockchain = "CBlockchain"
)

func (s *Storage) SaveCertBlock(block *types.CertBlock) error {
	_, err := s.Write(CBlockchain, block)
	if err != nil {
		return err
	}

	return nil
}
