package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes"
	"github.com/spacelavr/pandora/pkg/rpc"
	"github.com/spacelavr/pandora/pkg/storage"
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

	// todo when start node need auth with membership

	brkOpts, err := rpc.GetBrokerOpts(config.Viper.Tracker.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
	brk, err := broker.Connect(&broker.Opts{
		Endpoint: brkOpts.Endpoint,
		User:     brkOpts.User,
		Password: brkOpts.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	stg, err := storage.Connect(&storage.Opts{
		Endpoint: config.Viper.Database.Endpoint,
		Database: config.Viper.Database.Database,
		User:     config.Viper.Database.User,
		Password: config.Viper.Database.Password,
	})
	if err != nil {
		log.Fatal(err)
	}

	env.SetStorage(stg)
	env.SetRuntime(runtime.New())
	env.SetBroker(brk)

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := http.Listen(config.Viper.Node.Port, routes.Routes); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if config.Viper.Runtime.Clean {
			if err := stg.Clean(); err != nil {
				log.Error(err)
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
