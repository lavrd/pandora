package runtime

import (
	"time"

	"github.com/spacelavr/pandora/pkg/types"
)

type Runtime struct {
	blockchain types.MasterChain
}

func New() *Runtime {
	rt := &Runtime{}
	// rt.blockchain = types.Blockchain{rt.Genesis()}
	return rt
}

func (rt *Runtime) Add() error {
	return nil
}

// Genesis returns genesis block
func (_ *Runtime) Genesis() *types.Block {
	block := &types.Block{
		PrevHash:  "",
		Index:     0,
		Timestamp: time.Now().UTC(),
	}

	// block.Hash = sha256.SumString(block.Bytes())

	return block
}
