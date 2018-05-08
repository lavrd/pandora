package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
)

// Listen listen for events
func Listen() error {
	var (
		brk       = env.GetBroker()
		chrCBlock = make(chan *types.Block)
		chsNBlock = make(chan *types.Block)
	)

	if err := brk.Subscribe(broker.SCBlock, chrCBlock); err != nil {
		return err
	}

	if err := brk.Publish(broker.SNBlock, chsNBlock); err != nil {
		return err
	}

	for {
		select {
		case block, ok := <-chrCBlock:
			if !ok {
				return nil
			}

			chsNBlock <- block
		}
	}
}
