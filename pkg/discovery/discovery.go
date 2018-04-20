package discovery

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/discovery/env"
	"github.com/spacelavr/pandora/pkg/discovery/events"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/viper"
)

// Daemon start discovery daemon
func Daemon() bool {
	log.Debug("start discovery daemon")

	var (
		sig = make(chan os.Signal, 1)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	brk, err := broker.Connect(viper.GetString("broker.endpoint"))
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

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
