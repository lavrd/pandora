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
	)

	if err := env.GetBroker().Subscribe(broker.SubjectCertificate, chReadCert); err != nil {
		return err
	}

	if err := env.GetBroker().Publish(broker.SubjectBlock, chSendBlock); err != nil {
		return err
	}

	for {
		select {
		case cert := <-chReadCert:
			chSendBlock <- &types.Block{
				Cert: cert,
			}
		}
	}
}
