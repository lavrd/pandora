package node

import (
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

type Runtime struct {
	mc  *pb.MasterChain
	cc  *pb.CertChain
	lmb int
	lcb int

	// todo store at runtime or in another place (env?)
	FullName  string
	PublicKey string
}

func (r *Runtime) MC() *pb.MasterChain {
	return r.mc
}

func New(key *pb.PublicKey) (*Runtime, error) {
	mc, err := rpc.Node(key)
	if err != nil {
		return nil, err
	}

	cc := &pb.CertChain{}

	// todo for what with start?
	for _, mb := range mc.MasterChain {
		if mb.Block.PublicKey.PublicKey == key.PublicKey {
			cc = mb.CertChain
		}
	}

	return &Runtime{
		mc:  mc,
		cc:  cc,
		lcb: len(cc.CertChain) - 1,
		lmb: len(mc.MasterChain) - 1,
	}, nil
}
