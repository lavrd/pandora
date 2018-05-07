package events

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Listen listen for events
func Listen() error {
	var ()

	for {
		select {}
	}
}

func ReqSCAccount(opts *request.AccountCreate) *errors.Response {
	var (
		brk  = env.GetBroker()
		errR = &errors.Response{}
	)

	if err := brk.Request(broker.SCAccount, opts, errR); err != nil {
		return errors.InternalServerError()
	}

	if errR.Code != 0 {
		return errR
	}

	return nil
}

func ReqSFAccount(opts *request.AccountFetch) (*types.Account, *errors.Response) {
	var (
		brk = env.GetBroker()
		res = &types.ResAccountFetch{}
	)

	if err := brk.Request(broker.SFAccount, opts, res); err != nil {
		return nil, errors.InternalServerError()
	}

	log.Debug(res.Err)
	log.Debug(res.Acc)

	return res.Acc, res.Err
}
