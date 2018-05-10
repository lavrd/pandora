package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/types"
)

type Events struct {
	chsMasterBlock chan *types.MasterBlock
	chsCertBlock   chan *types.CertBlock
}

type Opts struct {
	*broker.Broker
}

func New() *Events {
	return &Events{
		chsMasterBlock: make(chan *types.MasterBlock),
		chsCertBlock:   make(chan *types.CertBlock),
	}
}

func (e *Events) Listen(opts *Opts) error {
	if err := opts.Broker.Publish(broker.SMasterBlock, e.chsMasterBlock); err != nil {
		return err
	}

	if err := opts.Broker.Publish(broker.SCertBlock, e.chsCertBlock); err != nil {
		return err
	}

	for {
		select {}
	}
}

func (e *Events) PMasterBlock(block *types.MasterBlock) {
	e.chsMasterBlock <- block
}

func (e *Events) PCertBlock(block *types.CertBlock) {
	e.chsCertBlock <- block
}
