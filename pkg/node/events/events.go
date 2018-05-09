package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Listen listen for events
func Listen() error {
	var (
		chrMasterBlock = make(chan *types.MasterBlock)
		brk            = env.GetBroker()
	)

	if err := brk.Subscribe(broker.SMasterBlock, chrMasterBlock); err != nil {
		return err
	}

	for {
		select {
		case block, ok := <-chrMasterBlock:
			if !ok {
				return nil
			}

			log.Debug(block)
			log.Debug(block.PublicKey)
		}
	}
}
