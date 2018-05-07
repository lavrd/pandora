package events

import (
	"encoding/json"

	"github.com/nats-io/go-nats"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/membership/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Listen listen for events
func Listen() error {
	if err := RepSCAccount(); err != nil {
		return err
	}

	if err := RepSFAccount(); err != nil {
		return err
	}

	for {
		select {}
	}
}

func RepSFAccount() error {
	var (
		brk = env.GetBroker()
		rt  = env.GetRuntime()
	)

	if err := brk.Reply(broker.SFAccount, func(m *nats.Msg) {
		var (
			opts = &request.AccountFetch{}
			errR *errors.Response
			acc  *types.Account
		)

		if err := json.Unmarshal(m.Data, opts); err != nil {
			log.Error(err)
			// todo this error will never happen
			errR = errors.InvalidJSON()
		} else {
			acc, err = rt.AccountFetch(opts)
			if err != nil {
				errR = errors.InternalServerError()
			}
			if acc == nil {
				errR = errors.NotFound("account")
			}
		}

		// todo return true or false
		if err := brk.SendReply(m.Reply, &types.ResAccountFetch{
			Acc: acc,
			Err: errR,
		}); err != nil {
			log.Error(err)
		}
	}); err != nil {
		return err
	}

	return nil
}

func RepSCAccount() error {
	var (
		brk = env.GetBroker()
		rt  = env.GetRuntime()
	)

	if err := brk.Reply(broker.SCAccount, func(m *nats.Msg) {
		var (
			opts = &request.AccountCreate{}
			errR *errors.Response
		)

		if err := json.Unmarshal(m.Data, opts); err != nil {
			log.Error(err)
			// todo this error will never happen
			errR = errors.InvalidJSON()
		} else {
			if err := rt.AccountCreate(opts); err != nil {
				if err == errors.AccountAlreadyExists {
					errR = errors.AlreadyExists("account")
				} else {
					errR = errors.InternalServerError()
				}
			}
		}

		// todo return true or false
		if err := brk.SendReply(m.Reply, errR); err != nil {
			log.Error(err)
		}
	}); err != nil {
		return err
	}

	return nil
}
