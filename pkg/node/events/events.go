package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/distribution"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/pb"
)

type events struct {
	chrMasterBlock chan *pb.MasterBlock
	chrCertBlock   chan *pb.CertBlock
	chrCert        chan *pb.Cert
}

func New(brk *broker.Broker) (*events, error) {
	var (
		chrCert        = make(chan *pb.Cert)
		chrMasterBlock = make(chan *pb.MasterBlock)
		chrCertBlock   = make(chan *pb.CertBlock)
	)

	if err := brk.Subscribe(broker.SubCertBlock, chrCertBlock); err != nil {
		return nil, err
	}

	if err := brk.Subscribe(broker.SubCert, chrCert); err != nil {
		return nil, err
	}

	if err := brk.Subscribe(broker.SubMasterBlock, chrMasterBlock); err != nil {
		return nil, err
	}

	return &events{
		chrCert:        chrCert,
		chrMasterBlock: chrMasterBlock,
		chrCertBlock:   chrCertBlock,
	}, nil
}

func (e *events) Listen() error {
	var (
		bc = env.GetBlockchain()
	)

	for {
		select {
		case cert, ok := <-e.chrCert:
			if !ok {
				return nil
			}

			if err := distribution.New().SaveCert(cert); err != nil {
				return err
			}
		case block, ok := <-e.chrMasterBlock:
			if !ok {
				return nil
			}

			bc.CommitMBlock(block)
		case block, ok := <-e.chrCertBlock:
			if !ok {
				return nil
			}

			bc.CommitCBlock(block)
		}
	}
}
