package master

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/master/events"
	"github.com/spacelavr/pandora/pkg/master/rpc"
	"github.com/spacelavr/pandora/pkg/master/runtime"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func Daemon() bool {
	log.Debug("start master daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	brk, err := broker.Connect(&broker.Opts{
		Endpoint: config.Viper.Discovery.Broker.Endpoint,
		User:     config.Viper.Discovery.Broker.User,
		Password: config.Viper.Discovery.Broker.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	env.SetBroker(brk)
	env.SetRuntime(runtime.New())
	env.SetEvents(events.New())

	go func() {
		if err := rpc.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := env.GetEvents().Listen(&events.Opts{Broker: brk}); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
