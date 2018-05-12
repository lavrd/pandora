package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/pb"
)

type Events struct {
	chsMasterBlock chan *pb.MasterBlock
	chsCertBlock   chan *pb.CertBlock
}

type Opts struct {
	*broker.Broker
}

func New() *Events {
	return &Events{
		chsMasterBlock: make(chan *pb.MasterBlock),
		chsCertBlock:   make(chan *pb.CertBlock),
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

func (e *Events) PMasterBlock(block *pb.MasterBlock) {
	e.chsMasterBlock <- block
}

func (e *Events) PCertBlock(block *pb.CertBlock) {
	e.chsCertBlock <- block
}
