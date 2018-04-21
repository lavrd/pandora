package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/api/env"
	"github.com/spacelavr/pandora/pkg/api/events"
	"github.com/spacelavr/pandora/pkg/api/routes"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spf13/viper"
)

// Daemon start api daemon
func Daemon() bool {
	log.Debug("start api daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	stg, err := storage.Connect(&storage.ConnectOpts{
		Endpoint: viper.GetString("db.endpoint"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer stg.Close()

	brk, err := broker.Connect(viper.GetString("broker.endpoint"))
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	env.SetStorage(stg)
	env.SetBroker(brk)

	go func() {
		if err := http.Listen(viper.GetInt("api.port"), routes.Routes); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if viper.GetBool("clean") {
			if err := stg.Clean(); err != nil {
				log.Error(err)
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
