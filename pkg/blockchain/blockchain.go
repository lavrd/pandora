package blockchain

import (
	"time"

	"pandora/pkg/pb"
	"pandora/pkg/utils/crypto/sha256"
)

// Blockchain
type Blockchain struct {
	mc *pb.MasterChain
	// last index master block
	limb int
}

// New returns new blockchain
func New() *Blockchain {
	bc := &Blockchain{limb: 0}
	bc.mc = &pb.MasterChain{MasterChain: []*pb.MasterBlock{bc.genesisMaster()}}
	return bc
}

// Sync sync blockchain
func Sync(mc *pb.MasterChain) *Blockchain {
	return &Blockchain{
		limb: len(mc.MasterChain) - 1,
		mc:   mc,
	}
}

// GetMasterChain returns master chain
func (bc *Blockchain) GetMasterChain() *pb.MasterChain {
	return bc.mc
}

// genesisMaster returns new master genesis block
func (bc *Blockchain) genesisMaster() *pb.MasterBlock {
	b := &pb.MasterBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: &pb.PublicKey{},
		},
		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{bc.genesisCert(nil)}},
	}

	b.Block.Hash = sha256.CalcToString(b.String())

	return b
}

// genesisCert returns new genesis cert
func (bc *Blockchain) genesisCert(key *pb.PublicKey) *pb.CertBlock {
	b := &pb.CertBlock{
		Block: &pb.Block{
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
	}

	b.Block.Hash = sha256.CalcToString(b.String())

	return b
}

// PrepareCertBlock prepare cert block
func (bc *Blockchain) PrepareCertBlock(cert *pb.Cert) *pb.CertBlock {
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
			Tx:        cert.ID,
		},
	}

	b.Block.Hash = sha256.CalcToString(b.String())

	return b
}

// PrepareMasterBlock prepare master block
func (bc *Blockchain) PrepareMasterBlock(key *pb.PublicKey) *pb.MasterBlock {
	b := &pb.MasterBlock{
		Block: &pb.Block{
			Index:     bc.mc.MasterChain[bc.limb].Block.Index + 1,
			PrevHash:  bc.mc.MasterChain[bc.limb].Block.Hash,
			Timestamp: time.Now().UTC().Unix(),
			PublicKey: key,
		},
		CertChain: &pb.CertChain{CertChain: []*pb.CertBlock{bc.genesisCert(key)}},
	}

	hash := sha256.CalcToString(b.String())
	b.Block.Hash = hash

	return b
}

// CommitMasterBlock commit master block
func (bc *Blockchain) CommitMasterBlock(b *pb.MasterBlock) {
	bc.mc.MasterChain = append(bc.mc.MasterChain, b)
	bc.limb++
}

// CommitCertBlock commit cert block
func (bc *Blockchain) CommitCertBlock(b *pb.CertBlock) {
	for i, mb := range bc.mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == b.Block.PublicKey.PublicKey {
			bc.mc.MasterChain[i].CertChain.CertChain = append(bc.mc.MasterChain[i].CertChain.CertChain, b)
			break
		}
	}
}
