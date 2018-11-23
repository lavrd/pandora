package distribution

import (
	"time"

	"pandora/pkg/blockchain"
	"pandora/pkg/node/env"
	"pandora/pkg/node/routes/request"
	"pandora/pkg/node/rpc"
	"pandora/pkg/pb"
	"pandora/pkg/storage/leveldb"
	"pandora/pkg/utils/generator"
)

// Distribution
type Distribution struct {
	storage *leveldb.Leveldb
	rpc     *rpc.RPC
	key     *pb.PublicKey
	bc      *blockchain.Blockchain
}

// New returns new distribution
func New() *Distribution {
	return &Distribution{
		storage: env.GetStorage(),
		key:     env.GetKey(),
		rpc:     env.GetRPC(),
		bc:      env.GetBlockchain(),
	}
}

// VerifyCert verify cert
func (d *Distribution) VerifyCert(opts *request.CertVerify) bool {
	for _, mb := range d.bc.GetMasterChain().MasterChain {
		for _, cb := range mb.CertChain.CertChain {
			if cb.Block.Tx == *opts.Id {
				return true
			}
		}
	}
	return false
}

// SaveCertBlock save cert block to storage
func (d *Distribution) SaveCertBlock(cb *pb.CertBlock) error {
	return d.storage.PutCertBlock(cb)
}

// SaveMasterBlock save master block to storage
func (d *Distribution) SaveMasterBlock(mb *pb.MasterBlock) error {
	return d.storage.PutMasterBlock(mb)
}

// SaveCert save cert to storage
func (d *Distribution) SaveCert(cert *pb.Cert) error {
	return d.storage.PutCert(cert)
}

// LoadCert load cert from storage
func (d *Distribution) LoadCert(id string) (*pb.Cert, error) {
	return d.storage.LoadCert(id)
}

// ProposeMember propose member
func (d *Distribution) ProposeMember(opts *request.Candidate) (*pb.PublicKey, error) {
	key, err := d.rpc.ProposeMember(&pb.MemberMeta{
		Name:  *opts.Name,
		Email: *opts.Email,
	})
	if err != nil {
		return nil, err
	}
	return key, nil
}

// FetchMember fetch member
func (d *Distribution) FetchMember(opts *request.MemberFetch) (*pb.Member, error) {
	mem, err := d.rpc.FetchMember(&pb.PublicKey{
		PublicKey: *opts.PublicKey,
	})
	if err != nil {
		return nil, err
	}
	return mem, nil
}

// SignCert sign cert
func (d *Distribution) SignCert(opts *request.CertIssue) error {
	return d.rpc.SignCert(&pb.Cert{
		ID: generator.ID(),
		Meta: &pb.CertMeta{
			Timestamp:   time.Now().UTC().UnixNano() / 1000000,
			Description: *opts.Description,
			Title:       *opts.Title,
		},
		Recipient: &pb.Participant{
			PublicKey: &pb.PublicKey{
				PublicKey: *opts.PublicKey,
			},
		},
		Issuer: &pb.Participant{
			PublicKey: &pb.PublicKey{
				PublicKey: d.key.PublicKey,
			},
		},
	})
}
