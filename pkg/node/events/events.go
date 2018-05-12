package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage"
)

// Listen listen for events
func Listen() error {
	var (
		chrMasterBlock = make(chan *pb.MasterBlock)
		chrCertBlock   = make(chan *pb.CertBlock)
		chrCert        = make(chan *pb.Cert)
		brk            = env.GetBroker()
		rt             = env.GetRuntime()
	)

	if err := brk.Subscribe(broker.SubCB, chrCertBlock); err != nil {
		return err
	}

	if err := brk.Subscribe(broker.SubCert, chrCert); err != nil {
		return err
	}

	if err := brk.Subscribe(broker.SubMB, chrMasterBlock); err != nil {
		return err
	}

	for {
		select {
		case cert, ok := <-chrCert:
			if !ok {
				return nil
			}

			if err := storage.Put(cert.Id, cert); err != nil {
				return err
			}
		case block, ok := <-chrMasterBlock:
			if !ok {
				return nil
			}

			rt.AddMC(block)
		case block, ok := <-chrCertBlock:
			if !ok {
				return nil
			}

			rt.AddCC(block)
		}
	}
}
