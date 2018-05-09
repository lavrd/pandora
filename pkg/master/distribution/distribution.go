package distribution

import (
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/pb"
)

func AddMasterBlock(publicKey *pb.PublicKey) {
	var (
		rt = env.GetRuntime()
		e  = env.GetEvents()
	)

	e.PMasterBlock(rt.AddMasterBlock(publicKey))
}
