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

func New(brk *broker.Broker) (*Events, error) {
	var (
		chsCert        = make(chan *pb.Cert)
		chsMasterBlock = make(chan *pb.MasterBlock)
		chsCertBlock   = make(chan *pb.CertBlock)
	)

	if err := brk.Publish(broker.SUB_MASTER_BLOCK, chsMasterBlock); err != nil {
		return nil, err
	}

	if err := brk.Publish(broker.SUB_CERT_BLOCK, chsCertBlock); err != nil {
		return nil, err
	}

	if err := brk.Publish(broker.SUB_CERT, chsCert); err != nil {
		return nil, err
	}

	return &Events{
		chsCert:        chsCert,
		chsMasterBlock: chsMasterBlock,
		chsCertBlock:   chsCertBlock,
	}, nil
}

func (e *Events) PubMasterBlock(block *pb.MasterBlock) {
	e.chsMasterBlock <- block
}

func (e *Events) PubCert(cert *pb.Cert) {
	e.chsCert <- cert
}

func (e *Events) PubCertBlock(block *pb.CertBlock) {
	e.chsCertBlock <- block
}
