package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chBlock = make(chan *types.Block)
	)

	err := env.GetBroker().Subscribe(broker.SubjectBlock, chBlock)
	if err != nil {
		return err
	}

	for {
		select {
		case _, ok := <-chBlock:
			if !ok {
				return nil
			}
		}
	}

	return nil
}
