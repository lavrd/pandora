package core

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/core/env"
	"github.com/spacelavr/pandora/pkg/core/events"
	"github.com/spacelavr/pandora/pkg/core/runtime"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/viper"
)

// Daemon start core daemon
func Daemon() bool {
	log.Debug("start core daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	brk, err := broker.Connect(viper.GetString("broker.endpoint"))
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	env.SetRuntime(&runtime.Runtime{})
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
