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
}

func New(key *pb.PublicKey) (*Runtime, error) {
	mc, err := rpc.Node(key)
	if err != nil {
		return nil, err
	}

	cc := &pb.CertChain{}

	for _, mb := range mc.Master_Block {
		if mb.PublicKey.PublicKey == key.PublicKey {
			cc = mb.CertChain
		}
	}

	return &Runtime{
		MC:  converter.FPBMC(mc),
		CC:  converter.FPBCC(cc),
		LCB: len(cc.CertBlock) - 1,
		LMB: len(mc.Master_Block) - 1,
	}, nil
}
