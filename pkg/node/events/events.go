package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
)

// Listen listen for events
func Listen() error {
	var (
		chsAccount = make(chan *request.Account)
		brk        = env.GetBroker()
	)

	if err := brk.Publish(broker.SAccount, chsAccount); err != nil {
		return err
	}

	for {
		select {
		case acc, ok := <-env.ReadAccount():
			if !ok {
				return nil
			}

			chsAccount <- acc
		}
	}
}
