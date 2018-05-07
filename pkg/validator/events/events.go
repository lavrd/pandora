package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/validator/env"
)

// Listen listen for events
func Listen() error {
	var (
		chReadNewBlock = make(chan *types.Block)
		chSendNewBlock = make(chan *types.Block)
		brk            = env.GetBroker()
	)

	if err := brk.Subscribe(broker.SBlock, chReadNewBlock); err != nil {
		return err
	}

	if err := brk.Publish(broker.SBlock, chSendNewBlock); err != nil {
		return err
	}

	for {
		select {
		case block, ok := <-chReadNewBlock:
			if !ok {
				return nil
			}

			chSendNewBlock <- block
		}
	}
}
