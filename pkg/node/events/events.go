package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chrMasterBlock = make(chan *types.MasterBlock)
		chrCertBlock   = make(chan *types.CertBlock)
		brk            = env.GetBroker()
		rt             = env.GetRuntime()
	)

	if err := brk.Subscribe(broker.SCertBlock, chrCertBlock); err != nil {
		return err
	}

	if err := brk.Subscribe(broker.SMasterBlock, chrMasterBlock); err != nil {
		return err
	}

	for {
		select {
		case block, ok := <-chrMasterBlock:
			if !ok {
				return nil
			}

			rt.AddMC(block)
		case block, ok := <-chrCertBlock:
			if !ok {
				return nil
			}

			rt.AddCC(block)
		}
	}
}
