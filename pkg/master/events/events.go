package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/types"
)

type Events struct {
	chsMasterBlock chan *types.MasterBlock
}

type Opts struct {
	*broker.Broker
}

func New() *Events {
	return &Events{
		chsMasterBlock: make(chan *types.MasterBlock),
	}
}

func (e *Events) Listen(opts *Opts) error {
	if err := opts.Broker.Publish(broker.SMasterBlock, e.chsMasterBlock); err != nil {
		return err
	}

	for {
		select {}
	}
}

func (e *Events) PMasterBlock(block *types.MasterBlock) {
	e.chsMasterBlock <- block
}
