package validator

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/validator/env"
	"github.com/spacelavr/pandora/pkg/validator/events"
	"github.com/spacelavr/pandora/pkg/validator/rpc"
	"github.com/spacelavr/pandora/pkg/validator/runtime"
)

func Daemon() bool {
	log.Debug("start validator daemon")

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

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := rpc.NotifyTracker(brkOpts); err != nil {
		log.Fatal(err)
	}

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
