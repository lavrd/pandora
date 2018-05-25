package blockchain

import (
	"time"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/crypto"
)

type Blockchain struct {
	mc   *pb.MasterChain
	limb int64
}

func New() *Blockchain {
	chain := &Blockchain{}
	chain.mc = &pb.MasterChain{MasterChain: []*pb.MasterBlock{b.GenesisMaster()}}
	return chain
}

func (b *Blockchain) GetMC() *pb.MasterChain {
	return b.mc
}

func (b *Blockchain) GenesisMaster() *pb.MasterBlock {
	block := &pb.MasterBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
		},

		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{b.GenesisCert(nil)}},
	}

	hash := crypto.SumString(block.String())
	block.Block.Hash = hash

	return block
}

func (b *Blockchain) GenesisCert(key *pb.PublicKey) *pb.CertBlock {
	block := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
	}

	hash := crypto.SumString(block.String())
	block.Block.Hash = hash

	return block
}

func (b *Blockchain) PrepareCertBlock(cert *pb.Cert) *pb.CertBlock {
	old := &pb.CertBlock{}

	for _, mb := range b.mc.MasterChain {
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

	hash := crypto.SumString(block.String())
	block.Block.Hash = hash

	return block
}

func (b *Blockchain) AddMasterBlock(key *pb.PublicKey) {
	isFound := false
	for _, b := range b.mc.MasterChain {
		if b.Block.PublicKey.PublicKey == key.PublicKey {
			isFound = true
		}
	}

	if !isFound {
		block := &pb.MasterBlock{
			Block: &pb.Block{
				Index:     b.mc.MasterChain[b.limb].Block.Index + 1,
				PrevHash:  b.mc.MasterChain[b.limb].Block.Hash,
				Timestamp: time.Now().UTC().Unix(),
				PublicKey: key,
			},
			CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{b.GenesisCert(key)}},
		}

		hash := crypto.SumString(block.String())
		block.Block.Hash = hash

		b.mc.MasterChain = append(b.mc.MasterChain, block)
	}
}
