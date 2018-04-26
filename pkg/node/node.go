package node

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/runtime"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start node daemon
func Daemon() bool {
	log.Debug("start node daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	brk, err := broker.Connect(&broker.Opts{
		Endpoint: config.Viper.Broker.Endpoint,
		User:     config.Viper.Broker.User,
		Password: config.Viper.Broker.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	rt := runtime.New()

	env.SetRuntime(rt)
	env.SetBroker(brk)

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
