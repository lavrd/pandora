package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/distribution"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/routes"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/storage/leveldb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start node daemon
func Daemon() bool {
	log.Debug("start node daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	candidate := &request.Candidate{
		Name:  &config.Viper.Node.Meta.Name,
		Email: &config.Viper.Node.Meta.Email,
	}
	if err := candidate.Validate(); err != nil {
		log.Fatal(err.Message)
	}

	r, err := rpc.New()
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	env.SetRPC(r)

	key, err := distribution.New().ProposeMember(candidate)
	if err != nil {
		if err != errors.AlreadyExists {
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

	stg, err := leveldb.New(config.Viper.Node.Database.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		stg.Close()

		if config.Viper.Runtime.Clean {
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
		if err := http.Listen(config.Viper.Node.Endpoint, routes.Routes, "./dashboard/static/"); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
