package events

import (
	"pandora/pkg/broker"
	"pandora/pkg/pb"
)

// Events
type Events struct {
	chsMasterBlock chan *pb.MasterBlock
	chsCertBlock   chan *pb.CertBlock
	chsCert        chan *pb.Cert
}

// New returns new events
func New(brk *broker.Broker) (*Events, error) {
	chsCert := make(chan *pb.Cert)
	chsMasterBlock := make(chan *pb.MasterBlock)
	chsCertBlock := make(chan *pb.CertBlock)

	if err := brk.Publish(broker.SubMasterBlock, chsMasterBlock); err != nil {
		return nil, err
	}

	if err := brk.Publish(broker.SubCertBlock, chsCertBlock); err != nil {
		return nil, err
	}

	if err := brk.Publish(broker.SubCert, chsCert); err != nil {
		return nil, err
	}

	return &Events{
		chsCert:        chsCert,
		chsMasterBlock: chsMasterBlock,
		chsCertBlock:   chsCertBlock,
	}, nil
}

// PubMasterBlock publish master block
func (e *Events) PubMasterBlock(block *pb.MasterBlock) {
	e.chsMasterBlock <- block
}

// PubCert publish cert
func (e *Events) PubCert(cert *pb.Cert) {
	e.chsCert <- cert
}

// PubCertBlock publish cert block
func (e *Events) PubCertBlock(block *pb.CertBlock) {
	e.chsCertBlock <- block
}
