package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chReadNewBlock = make(chan *types.Block)
		brk            = env.GetBroker()
		rt             = env.GetRuntime()
	)

	if err := brk.Subscribe(broker.SubjectNewBlock, chReadNewBlock); err != nil {
		return err
	}

	if err := brk.Reply(broker.SubjectLastBlock, func() (interface{}, error) {
		return rt.LastBlock(), nil
	}); err != nil {
		return err
	}

	if err := brk.Reply(broker.SubjectBlockchain, func() (interface{}, error) {
		return rt.Blockchain(), nil
	}); err != nil {
		return err
	}

	for {
		select {
		case block, ok := <-chReadNewBlock:
			if !ok {
				return nil
			}

			rt.AddBlock(block)
		}
	}
}
