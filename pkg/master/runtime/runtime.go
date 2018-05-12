package runtime

import (
	"time"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// todo need optimization and reuse

type Runtime struct {
	limb int
	mc   *pb.MasterChain
}

func New() *Runtime {
	rt := &Runtime{}

	rt.mc = &pb.MasterChain{MasterChain: []*pb.MasterBlock{rt.GenesisMasterBlock()}}

	return rt
}

func (rt *Runtime) MC() *pb.MasterChain {
	return rt.mc
}

func (rt *Runtime) LastMasterBlock() *pb.MasterBlock {
	return rt.mc.MasterChain[rt.limb]
}

func (rt *Runtime) AddMasterBlock(publicKey *pb.PublicKey) *pb.MasterBlock {
	isFound := false
	for _, b := range rt.mc.MasterChain {
		if b.Block.PublicKey.PublicKey == publicKey.PublicKey {
			log.Debug("ISFOUND")
			isFound = true
		}
	}

	if !isFound {
		block := &pb.MasterBlock{
			Block: &pb.Block{
				Index:     rt.LastMasterBlock().Block.Index + 1,
				PrevHash:  rt.LastMasterBlock().Block.Hash,
				Timestamp: time.Now().UTC().Unix(),
				PublicKey: publicKey,
			},
			CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{rt.GenesisCertBlock(publicKey)}},
		}

		hash := sha256.SumString(block.String())

		block.Block.Hash = hash

		rt.mc.MasterChain = append(rt.mc.MasterChain, block)

		return block
	}

	return nil
}

func (rt *Runtime) PrepareBlock(cert *pb.Cert) *pb.CertBlock {
	old := &pb.CertBlock{}

	for _, mb := range rt.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == cert.Issuer.PublicKey.PublicKey {
			old = mb.CertChain.CertChain[len(mb.CertChain.CertChain)-1]
		}
	}

	block := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PrevHash:  old.Block.Hash,
			Index:     old.Block.Index + 1,
			PublicKey: old.Block.PublicKey,
			Tx:        cert.Id,
		},
	}

	hash := sha256.SumString(block.String())
	block.Block.Hash = hash

	return block
}

func (_ *Runtime) GenesisCertBlock(publicKey *pb.PublicKey) *pb.CertBlock {
	block := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: publicKey,
		},
	}

	hash := sha256.SumString(block.String())

	block.Block.Hash = hash

	return block
}

func (rt *Runtime) GenesisMasterBlock() *pb.MasterBlock {
	block := &pb.MasterBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: &pb.PublicKey{
				PublicKey: "",
			},
		},

		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{rt.GenesisCertBlock(nil)}},
	}

	hash := sha256.SumString(block.String())

	block.Block.Hash = hash

	return block
}
