package runtime

import (
	"time"

	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
)

// Runtime
type Runtime struct{}

// PrepareBlock prepare block
func (r *Runtime) PrepareBlock(cert *types.Certificate, last *types.Block) *types.Block {
	block := &types.Block{
		Cert:      cert,
		PrevHash:  last.Hash,
		Index:     last.Index + 1,
		Timestamp: time.Now().UTC(),
	}

	block.Hash = sha256.Compute(block.String())

	return block
}
