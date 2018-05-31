package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage/leveldb"
	"github.com/spacelavr/pandora/pkg/utils/generator"
)

type Distribution struct {
	storage *leveldb.LevelDB
	rpc     *rpc.RPC
	key     *pb.PublicKey
	bc      *blockchain.Blockchain
}

func New() *Distribution {
	return &Distribution{
		storage: env.GetStorage(),
		key:     env.GetKey(),
		rpc:     env.GetRPC(),
		bc:      env.GetBlockchain(),
	}
}

func (d *Distribution) VerifyCert(opts *request.CertVerify) bool {
	for _, mb := range d.bc.MasterChain().MasterChain {
		for _, cb := range mb.CertChain.CertChain {
			if cb.Block.Tx == *opts.Id {
				return true
			}
		}
	}
	return false
}

func (d *Distribution) SaveCert(cert *pb.Cert) error {
	return d.storage.Put(cert)
}

func (d *Distribution) LoadCert(id string) (*pb.Cert, error) {
	return d.storage.Load(id)
}

func (d *Distribution) MasterChain() *pb.MasterChain {
	return env.GetBlockchain().MasterChain()
}

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

func (d *Distribution) FetchMember(opts *request.MemberFetch) (*pb.Member, error) {
	mem, err := d.rpc.FetchMember(&pb.PublicKey{
		PublicKey: *opts.PublicKey,
	})
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (d *Distribution) SignCert(opts *request.CertIssue) error {
	return d.rpc.SignCert(&pb.Cert{
		Id: generator.Id(),
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
