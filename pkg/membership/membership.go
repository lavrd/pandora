package membership

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/membership/env"
	"github.com/spacelavr/pandora/pkg/membership/events"
	"github.com/spacelavr/pandora/pkg/membership/runtime"
	"github.com/spacelavr/pandora/pkg/rpc"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func Daemon() bool {
	log.Debug("start membership daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

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

	env.SetBroker(brk)
	env.SetRuntime(runtime.New(&runtime.Opts{Storage: stg}))

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
