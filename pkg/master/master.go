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

	brkOpts := &broker.Opts{
		Endpoint: config.Viper.Broker.Endpoint,
		User:     config.Viper.Broker.User,
		Password: config.Viper.Broker.Password,
	}
	brk, err := broker.Connect(brkOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	rt := runtime.New()

	env.SetBroker(brk)
	env.SetRuntime(rt)
	env.SetEvents(events.New(brk))

	go func() {
		if err := rpc.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := env.GetEvents().Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
