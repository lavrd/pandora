package runtime

import (
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/converter"
)

type Runtime struct {
	MC  types.MasterChain
	CC  types.CertChain
	LMB int
	LCB int

	// todo store at runtime or in another place (env?)
	FullName  string
	PublicKey string
}

func New(key *pb.PublicKey) (*Runtime, error) {
	mc, err := rpc.Node(key)
	if err != nil {
		return nil, err
	}

	cc := &pb.CertChain{}

	for _, mb := range mc.MasterBlock {
		if mb.PublicKey.PublicKey == key.PublicKey {
			cc = mb.CertChain
		}
	}

	return &Runtime{
		MC:  converter.FPBMC(mc),
		CC:  converter.FPBCC(cc),
		LCB: len(cc.CertBlock) - 1,
		LMB: len(mc.MasterBlock) - 1,
	}, nil
}

func (r *Runtime) AddMC(block *types.MasterBlock) {
	r.MC = append(r.MC, block)
}

func (r *Runtime) AddCC(block *types.CertBlock) {
	for i, mb := range r.MC {
		if mb.PublicKey == block.PublicKey {
			r.MC[i].CertChain = append(r.MC[i].CertChain, block)
		}
	}
}
