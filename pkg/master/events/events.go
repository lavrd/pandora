package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/types"
)

type Events struct {
	// todo :( not good
	*broker.Broker
	ChSMasterBlock chan *types.MasterBlock
}

func New(brk *broker.Broker) *Events {
	return &Events{
		Broker:         brk,
		ChSMasterBlock: make(chan *types.MasterBlock),
	}
}

// Listen listen for events
func (e *Events) Listen() error {
	if err := e.Publish(broker.SMasterBlock, e.ChSMasterBlock); err != nil {
		return err
	}

	for {
		select {}
	}
}

func (e *Events) PMasterBlock(block *types.MasterBlock) {
	e.ChSMasterBlock <- block
}
