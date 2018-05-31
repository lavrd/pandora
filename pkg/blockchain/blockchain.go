package blockchain

import (
	"time"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
)

type Blockchain struct {
	mc   *pb.MasterChain
	limb int
}

func New() *Blockchain {
	bc := &Blockchain{limb: 0}
	bc.mc = &pb.MasterChain{MasterChain: []*pb.MasterBlock{bc.genesisM()}}
	return bc
}

func Sync(mc *pb.MasterChain) *Blockchain {
	return &Blockchain{
		limb: len(mc.MasterChain) - 1,
		mc:   mc,
	}
}

func (bc *Blockchain) MasterChain() *pb.MasterChain {
	return bc.mc
}

func (bc *Blockchain) genesisM() *pb.MasterBlock {
	b := &pb.MasterBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: &pb.PublicKey{},
		},
		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{bc.genesisC(nil)}},
	}

	hash := sha256.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) genesisC(key *pb.PublicKey) *pb.CertBlock {
	b := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
	}

	hash := sha256.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) PrepareCBlock(cert *pb.Cert) *pb.CertBlock {
	old := &pb.CertBlock{}
	for _, mb := range bc.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == cert.Issuer.PublicKey.PublicKey {
			old = mb.CertChain.CertChain[len(mb.CertChain.CertChain)-1]
			break
		}
	}

	b := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PrevHash:  old.Block.Hash,
			Index:     old.Block.Index + 1,
			PublicKey: old.Block.PublicKey,
			Tx:        cert.Id,
		},
	}

	hash := sha256.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) PrepareMBlock(key *pb.PublicKey) *pb.MasterBlock {
	b := &pb.MasterBlock{
		Block: &pb.Block{
			Index:     bc.mc.MasterChain[bc.limb].Block.Index + 1,
			PrevHash:  bc.mc.MasterChain[bc.limb].Block.Hash,
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{bc.genesisC(key)}},
	}

	hash := sha256.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) CommitMBlock(b *pb.MasterBlock) {
	bc.mc.MasterChain = append(bc.mc.MasterChain, b)
	bc.limb++
}

func (bc *Blockchain) CommitCBlock(b *pb.CertBlock) {
	for i, mb := range bc.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == b.Block.PublicKey.PublicKey {
			bc.mc.MasterChain[i].CertChain.CertChain = append(bc.mc.MasterChain[i].CertChain.CertChain, b)
			break
		}
	}
}
