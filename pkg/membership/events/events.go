package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/membership/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Listen listen for events
func Listen() error {
	var (
		brk = env.GetBroker()
		rt  = env.GetRuntime()
	)

	if err := brk.Reply(broker.SNAccount, func(subject, reply string, msg interface{}) {
		if err := rt.AccountCreate(msg.(*request.Account)); err != nil {
			if err := brk.SendReply(reply, "bad"); err != nil {
				log.Error(err)
			}
		}
		if err := brk.SendReply(reply, "ok"); err != nil {
			log.Error(err)
		}
	}); err != nil {
		return err
	}

	for {
		select {}
	}
}
