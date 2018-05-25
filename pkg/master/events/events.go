package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/pb"
)

type Events struct {
	chsMasterBlock chan *pb.MasterBlock
	chsCertBlock   chan *pb.CertBlock
	chsCert        chan *pb.Cert
}

func New() *Events {
	return &Events{
		chsCert:        make(chan *pb.Cert),
		chsMasterBlock: make(chan *pb.MasterBlock),
		chsCertBlock:   make(chan *pb.CertBlock),
	}
}

func (e *Events) Listen(brk *broker.Broker) error {
	if err := brk.Publish(broker.SubMB, e.chsMasterBlock); err != nil {
		return err
	}

	if err := brk.Publish(broker.SubCB, e.chsCertBlock); err != nil {
		return err
	}

	if err := brk.Publish(broker.SubCert, e.chsCert); err != nil {
		return err
	}

	for {
		select {}
	}
}

func (e *Events) PMasterBlock(block *pb.MasterBlock) {
	e.chsMasterBlock <- block
}

func (e *Events) PubCert(cert *pb.Cert) {
	e.chsCert <- cert
}

func (e *Events) PCertBlock(block *pb.CertBlock) {
	e.chsCertBlock <- block
}
