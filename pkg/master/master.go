package master

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/master/events"
	"github.com/spacelavr/pandora/pkg/master/rpc"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func Daemon() bool {
	log.Debug("start master daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// todo master need get broker opts from discovery
	brk, err := broker.Connect(
		config.Viper.Discovery.Broker.Endpoint,
		config.Viper.Discovery.Broker.User,
		config.Viper.Discovery.Broker.Password,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	go func() {
		if err := rpc.New().Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	evt, err := events.New(brk)
	if err != nil {
		log.Fatal(err)
	}

	env.SetBlockchain(blockchain.New())
	env.SetEvents(evt)

	defer func() {
		if config.Viper.Runtime.Clean {
			// todo clean database
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
