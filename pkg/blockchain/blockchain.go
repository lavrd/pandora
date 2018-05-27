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
	bc := &Blockchain{}
	bc.mc = &pb.MasterChain{MasterChain: []*pb.MasterBlock{bc.GenesisMaster()}}
	return bc
}

func (bc *Blockchain) GetMasterChain() *pb.MasterChain {
	return bc.mc
}

func (bc *Blockchain) GenesisMaster() *pb.MasterBlock {
	b := &pb.MasterBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
		},
		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{bc.GenesisCert(nil)}},
	}

	hash := crypto.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) GenesisCert(key *pb.PublicKey) *pb.CertBlock {
	b := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
	}

	hash := crypto.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) PrepareCertBlock(cert *pb.Cert) *pb.CertBlock {
	old := &pb.CertBlock{}
	for _, mb := range bc.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == cert.Issuer.PublicKey.PublicKey {
			old = mb.CertChain.CertChain[len(mb.CertChain.CertChain)-1]
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

	hash := crypto.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) PrepareMasterBlock(key *pb.PublicKey) *pb.MasterBlock {
	b := &pb.MasterBlock{
		Block: &pb.Block{
			Index:     bc.mc.MasterChain[bc.limb].Block.Index + 1,
			PrevHash:  bc.mc.MasterChain[bc.limb].Block.Hash,
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{bc.GenesisCert(key)}},
	}

	hash := crypto.SumString(b.String())
	b.Block.Hash = hash

	return b
}

func (bc *Blockchain) CommitMasterBlock(b *pb.MasterBlock) {
	for _, mb := range bc.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == b.Block.PublicKey.PublicKey {
			bc.mc.MasterChain = append(bc.mc.MasterChain, b)
		}
	}
}

func (bc *Blockchain) CommitCertBlock(b *pb.CertBlock) {
	for i, mb := range bc.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == b.Block.PublicKey.PublicKey {
			bc.mc.MasterChain[i].CertChain.CertChain = append(bc.mc.MasterChain[i].CertChain.CertChain, b)
		}
	}
}
