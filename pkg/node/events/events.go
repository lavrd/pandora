package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
)

var (
	chsCBlock = make(chan *types.Block)
)

// Listen listen for events
func Listen() error {
	var (
		brk = env.GetBroker()
		// rt        = env.GetRuntime()
		chrNBlock = make(chan *types.Block)
	)

	if err := brk.Publish(broker.SCBlock, chsCBlock); err != nil {
		return err
	}

	if err := brk.QSubscribe(broker.SNBlock, broker.QCBlock, chrNBlock); err != nil {
		return err
	}

	for {
		select {}
	}
}
