package distribution

import (
	"time"

	"github.com/satori/go.uuid"

	"pandora/pkg/blockchain"
	"pandora/pkg/node/env"
	"pandora/pkg/node/routes/request"
	"pandora/pkg/node/rpc"
	"pandora/pkg/pb"
	"pandora/pkg/storage/leveldb"
)

// Node
type Node struct {
	storage *leveldb.Leveldb
	rpc     *rpc.RPC
	key     *pb.PublicKey
	bc      *blockchain.Blockchain
}

// VerifyCert verify cert
func (n *Node) VerifyCert(opts *request.CertVerify) bool {
	for _, mb := range n.bc.GetMasterChain().MasterChain {
		for _, cb := range mb.CertChain.CertChain {
			if cb.Block.Tx == *opts.Id {
				return true
			}
		}
	}
	return false
}

// SaveCertBlock save cert block to storage
func (n *Node) SaveCertBlock(cb *pb.CertBlock) error {
	return n.storage.PutCertBlock(cb)
}

// SaveMasterBlock save master block to storage
func (n *Node) SaveMasterBlock(mb *pb.MasterBlock) error {
	return n.storage.PutMasterBlock(mb)
}

// SaveCert save cert to storage
func (n *Node) SaveCert(cert *pb.Cert) error {
	return n.storage.PutCert(cert)
}

// LoadCert load cert from storage
func (n *Node) LoadCert(id string) (*pb.Cert, error) {
	return n.storage.LoadCert(id)
}

// ProposeMember propose member
func (n *Node) ProposeMember(opts *request.Candidate) (*pb.PublicKey, error) {
	key, err := n.rpc.ProposeMember(&pb.MemberMeta{
		Name:  *opts.Name,
		Email: *opts.Email,
	})
	if err != nil {
		return nil, err
	}
	return key, nil
}

// FetchMember fetch member
func (n *Node) FetchMember(opts *request.MemberFetch) (*pb.Member, error) {
	mem, err := n.rpc.FetchMember(&pb.PublicKey{
		PublicKey: *opts.PublicKey,
	})
	if err != nil {
		return nil, err
	}
	return mem, nil
}

// SignCert sign cert
func (n *Node) SignCert(opts *request.CertIssue) error {
	return n.rpc.SignCert(&pb.Cert{
		ID: uuid.NewV1().String(),
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
				PublicKey: n.key.PublicKey,
			},
		},
	})
}

// New returns new node distribution
func NewNode() *Node {
	return &Node{
		storage: env.GetStorage(),
		key:     env.GetKey(),
		rpc:     env.GetRPC(),
		bc:      env.GetBlockchain(),
	}
}
