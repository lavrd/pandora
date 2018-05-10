package distribution

import (
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/converter"
)

func AddMasterBlock(publicKey *pb.PublicKey) {
	var (
		rt = env.GetRuntime()
		e  = env.GetEvents()
	)

	e.PMasterBlock(rt.AddMasterBlock(publicKey))
}

func GetMasterChain() *pb.MasterChain {
	var (
		rt = env.GetRuntime()
	)

	return converter.TPBMC(rt.MC)
}
