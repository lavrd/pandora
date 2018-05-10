package distribution

import (
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
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

func PrepareBlock(cert *pb.Cert) *types.CertBlock {
	var (
		rt = env.GetRuntime()
	)

	return rt.PrepareBlock(cert)
}
