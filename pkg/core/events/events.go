package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/core/env"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chReadNewCert  = make(chan *types.Certificate)
		chSendNewBlock = make(chan *types.Block)
		brk            = env.GetBroker()
		rt             = env.GetRuntime()
	)

	if err := brk.Subscribe(broker.SubjectNewCertificate, chReadNewCert); err != nil {
		return err
	}

	if err := brk.Publish(broker.SubjectNewBlock, chSendNewBlock); err != nil {
		return err
	}

	for {
		select {
		case cert, ok := <-chReadNewCert:
			if !ok {
				return nil
			}

			last := &types.Block{}
			if err := brk.Request(broker.SubjectLastBlock, "", last); err != nil {
				return err
			}

			chSendNewBlock <- rt.PrepareBlock(cert, last)
		}
	}
}
