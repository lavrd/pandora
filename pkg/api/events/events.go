package events

import (
	"github.com/spacelavr/pandora/pkg/api/env"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/types"
)

// Listen listen for events
func Listen() error {
	var (
		chSendNewCert = make(chan *types.Certificate)
		brk           = env.GetBroker()
	)

	if err := brk.Publish(broker.SubjectNewCertificate, chSendNewCert); err != nil {
		return err
	}

	for {
		select {
		case cert, ok := <-env.ReadCert():
			if !ok {
				return nil
			}

			chSendNewCert <- cert
		}
	}
}
