package runtime

import (
	"time"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// todo need optimization and reuse

type Runtime struct {
	LIMB int
	MC   types.MasterChain
}

func New() *Runtime {
	rt := &Runtime{
		LIMB: 0,
	}

	rt.MC = types.MasterChain{rt.GenesisMasterBlock()}

	return rt
}

func (rt *Runtime) LastMasterBlock() *types.MasterBlock {
	return rt.MC[rt.LIMB]
}

func (rt *Runtime) AddMasterBlock(publicKey *pb.PublicKey) *types.MasterBlock {
	isFound := false
	for _, b := range rt.MC {
		if b.PublicKey == publicKey.PublicKey {
			log.Debug("ISFOUND")
			isFound = true
		}
	}

	if !isFound {
		block := &types.MasterBlock{
			PublicKey: publicKey.PublicKey,
			Block: &types.Block{
				Index:     rt.LastMasterBlock().Index + 1,
				PrevHash:  rt.LastMasterBlock().Hash,
				Timestamp: time.Now().UTC(),
			},
			CertChain: types.CertChain{rt.GenesisCertBlock()},
		}

		hash := sha256.SumString(block.Bytes())

		block.Hash = hash
		block.Key = hash

		rt.MC = append(rt.MC, block)

		return block
	}

	return nil
}

func (_ *Runtime) GenesisCertBlock() *types.CertBlock {
	block := &types.CertBlock{
		Block: &types.Block{
			Index:     0,
			PrevHash:  "",
			Timestamp: time.Now().UTC(),
		},
	}

	hash := sha256.SumString(block.Bytes())

	block.Hash = hash
	block.Key = hash

	return block
}

func (rt *Runtime) GenesisMasterBlock() *types.MasterBlock {
	block := &types.MasterBlock{
		Block: &types.Block{
			Index:     0,
			PrevHash:  "",
			Timestamp: time.Now().UTC(),
		},
		CertChain: types.CertChain{rt.GenesisCertBlock()},
	}

	hash := sha256.SumString(block.Bytes())

	block.Hash = hash
	block.Key = hash

	return block
}
