package node

import (
	"pandora/pkg/blockchain"
	"pandora/pkg/broker"
	"pandora/pkg/conf"
	"pandora/pkg/distribution"
	"pandora/pkg/node/env"
	"pandora/pkg/node/events"
	"pandora/pkg/node/routes"
	"pandora/pkg/node/routes/request"
	"pandora/pkg/node/rpc"
	"pandora/pkg/storage/leveldb"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/http"
	"pandora/pkg/utils/log"
	"pandora/pkg/utils/signalutils"
)

const (
	NODE = "node"
)

// Daemon start node daemon
func Daemon() bool {
	log.Debug("start node daemon")

	candidate := &request.Candidate{
		Name:  &conf.Conf.Node.Meta.Name,
		Email: &conf.Conf.Node.Meta.Email,
	}
	if err := candidate.Validate(); err != nil {
		log.Fatal(errors.New(err.Message))
	}

	r, err := rpc.New()
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	env.SetRPC(r)

	key, err := distribution.NewNode().ProposeMember(candidate)
	if err != nil {
		if err != errors.ErrAlreadyExists {
			log.Fatal(err)
		}
	}

	mc, brkOpts, err := r.InitNode(key)
	if err != nil {
		log.Error(err)
	}

	brk, err := broker.New(
		brkOpts.Endpoint,
		brkOpts.User,
		brkOpts.Password,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	stg, err := leveldb.New(conf.Conf.Node.Leveldb)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := stg.Close(); err != nil {
			log.Error(err)
		}

		if conf.Conf.Runtime.Clean {
			if err := stg.Clean(); err != nil {
				log.Error(err)
			}
		}
	}()

	evt, err := events.New(brk)
	if err != nil {
		log.Fatal(err)
	}

	env.SetBlockchain(blockchain.Sync(mc))
	env.SetStorage(stg)
	env.SetKey(key)

	go func() {
		if err := evt.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := http.Listen(conf.Conf.Node.Endpoint, routes.SubRoutes, "./dashboard/static/"); err != nil {
			log.Fatal(err)
		}
	}()

	signalutils.NotifyInterrupt()
	return true
}
