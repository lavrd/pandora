package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/api/env"
	"github.com/spacelavr/pandora/pkg/api/events"
	"github.com/spacelavr/pandora/pkg/api/routes"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start api daemon
func Daemon() bool {
	log.Debug("start api daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	stg, err := storage.Connect(&storage.Opts{
		Endpoint: config.Viper.Database.Endpoint,
		User:     config.Viper.Database.User,
		Password: config.Viper.Database.Password,
		Database: config.Viper.Database.Database,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer stg.Close()

	brk, err := broker.Connect(&broker.Opts{
		Endpoint: config.Viper.Broker.Endpoint,
		User:     config.Viper.Broker.User,
		Password: config.Viper.Broker.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	env.SetStorage(stg)
	env.SetBroker(brk)

	go func() {
		if err := http.Listen(config.Viper.Api.Port, routes.Routes); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := events.Listen(); err != nil {
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
