package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/types"
)

// Genesis generate genesis block
func (d *Distribution) Genesis() *types.Block {
	block := &types.Block{
		Certificate: nil,
		Index:       0,
		PrevHash:    "",
		Timestamp:   time.Now().UTC(),
	}

	// todo need calc sha256 hash
	block.Hash = block.String()

	return block
}

// Init init blockchain
func (d *Distribution) Init() types.Blockchain {
	return types.Blockchain{d.Genesis()}
}
