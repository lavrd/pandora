package events

import (
	"pandora/pkg/broker"
	"pandora/pkg/distribution"
	"pandora/pkg/node/env"
	"pandora/pkg/pb"
)

type events struct {
	chrMasterBlock chan *pb.MasterBlock
	chrCertBlock   chan *pb.CertBlock
	chrCert        chan *pb.Cert
}

// New returns new events
func New(brk *broker.Broker) (*events, error) {
	var (
		chrCert        = make(chan *pb.Cert)
		chrMasterBlock = make(chan *pb.MasterBlock)
		chrCertBlock   = make(chan *pb.CertBlock)
	)

	if err := brk.Subscribe(broker.SUB_CERT_BLOCK, chrCertBlock); err != nil {
		return nil, err
	}

	if err := brk.Subscribe(broker.SUB_CERT, chrCert); err != nil {
		return nil, err
	}

	if err := brk.Subscribe(broker.SUB_MASTER_BLOCK, chrMasterBlock); err != nil {
		return nil, err
	}

	return &events{
		chrCert:        chrCert,
		chrMasterBlock: chrMasterBlock,
		chrCertBlock:   chrCertBlock,
	}, nil
}

// Listen listen for events
func (e *events) Listen() error {
	var (
		bc   = env.GetBlockchain()
		dist = distribution.NewNode()
	)

	for {
		select {
		case cert, ok := <-e.chrCert:
			if !ok {
				return nil
			}

			if err := dist.SaveCert(cert); err != nil {
				return err
			}
		case block, ok := <-e.chrMasterBlock:
			if !ok {
				return nil
			}

			if err := dist.SaveMasterBlock(block); err != nil {
				return err
			}

			bc.CommitMasterBlock(block)
		case block, ok := <-e.chrCertBlock:
			if !ok {
				return nil
			}

			if err := dist.SaveCertBlock(block); err != nil {
				return err
			}

			bc.CommitCertBlock(block)
		}
	}
}
