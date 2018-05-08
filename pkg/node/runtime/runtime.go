package runtime

import (
	"time"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/types"
	ed255192 "github.com/spacelavr/pandora/pkg/utils/crypto/ed25519"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"golang.org/x/crypto/ed25519"
)

// Runtime
type Runtime struct {
	FullName   string
	PrivateKey ed25519.PrivateKey
	last       int
	blockchain types.Blockchain
}

// New returns new runtime
func New() *Runtime {
	_, privateKey := ed255192.GenerateKeys()

	r := &Runtime{
		last:       0,
		PrivateKey: privateKey,
		FullName:   config.Viper.Node.FullName,
	}

	r.blockchain = types.Blockchain{r.Genesis()}

	return r
}

// PrepareBlock prepare block
func (r *Runtime) PrepareBlock(cert *types.Certificate, last *types.Block) *types.Block {
	block := &types.Block{
		Cert:      cert,
		PrevHash:  last.Hash,
		Index:     last.Index + 1,
		Timestamp: time.Now().UTC(),
	}

	block.Hash = sha256.SumString(block.Bytes())

	return block
}

// Genesis returns genesis block
func (r *Runtime) Genesis() *types.Block {
	block := &types.Block{
		Cert:      nil,
		PrevHash:  "",
		Index:     0,
		Timestamp: time.Now().UTC(),
	}

	block.Hash = sha256.SumString(block.Bytes())

	return block
}

// Validate validate block
func (r *Runtime) Validate(block *types.Block) bool {
	return true
}

// VBlockchain validate blockchain
func (r *Runtime) VBlockchain() bool {
	return true
}

// Add add block to blockchain
func (r *Runtime) Add(block *types.Block) {
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
