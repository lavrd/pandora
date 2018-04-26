package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/core/env"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chReadCert  = make(chan *types.Certificate)
		chSendBlock = make(chan *types.Block)
		brk         = env.GetBroker()
		rt          = env.GetRuntime()
	)

	if err := brk.Subscribe(broker.SCertificate, chReadCert); err != nil {
		return err
	}

	if err := brk.Publish(broker.SBlock, chSendBlock); err != nil {
		return err
	}

	for {
		select {
		case cert, ok := <-chReadCert:
			if !ok {
				return nil
			}

			last := &types.Block{}
			if err := brk.Request(broker.SLBlock, "", last); err != nil {
				return err
			}

			chSendBlock <- rt.PrepareBlock(cert, last)
		}
	}
}
