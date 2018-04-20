package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chReadBlock = make(chan *types.Block)
	)

	if err := env.GetBroker().Subscribe(broker.SubjectBlock, chReadBlock); err != nil {
		return err
	}

	for {
		select {
		case _, ok := <-chReadBlock:
			if !ok {
				return nil
			}
		}
	}

	return nil
}
