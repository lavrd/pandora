package runtime

import (
	"time"

	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
)

// Runtime
type Runtime struct {
	last        int
	CertChain   types.CertChain
	MasterBlock *types.MasterBlock
}

// New returns new runtime
func New() *Runtime {
	r := &Runtime{
		last: 0,
	}

	r.blockchain = types.Blockchain{r.Genesis()}

	return r
}

// Genesis returns genesis block
func (_ *Runtime) Genesis() *types.Block {
	block := &types.Block{
		PrevHash:  "",
		Index:     0,
		Timestamp: time.Now().UTC(),
	}

	block.Hash = sha256.SumString(block.Bytes())

	return block
}

func (rt *Runtime) Genesis() *types.CertBlock {
	block := &types.CertBlock{

	}

	return block
}

// PrepareBlock prepare block
func (r *Runtime) PrepareBlock(cert *types.Certificate, last *types.CertBlock) *types.CertChain {
	block := &types.Block{
		Cert:      cert,
		PrevHash:  last.Hash,
		Index:     last.Index + 1,
		Timestamp: time.Now().UTC(),
	}

	block.Hash = sha256.SumString(block.Bytes())

	return nil
}

// Validate validate block
func (r *Runtime) Validate(block *types.CertChain) bool {
	return true
}

// VBlockchain validate blockchain
func (r *Runtime) VBlockchain() bool {
	return true
}

// Add add block to blockchain
func (r *Runtime) Add(block *types.CertChain) {
	if r.Validate(block) {
		r.blockchain = append(r.blockchain, block)
		r.last++
	}
}

// Last returns last blockchain block
func (r *Runtime) Last() *types.Block {
	return r.blockchain[r.last]
}

// Blockchain returns blockchain
func (r *Runtime) Blockchain() types.Blockchain {
	return r.blockchain
}
