package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/api/env"
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
		sig = make(chan os.Signal, 1)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	stg, err := storage.Open(viper.GetString("db.file"))
	if err != nil {
		log.Fatal(err)
	}

	brk, err := broker.Connect(viper.GetString("broker.url"), viper.GetInt("broker.port"))
	if err != nil {
		log.Fatal(err)
	}

	env.SetStorage(stg)
	env.SetBroker(brk)

	go func() {
		if err := http.Listen(
			viper.GetString("api.host"),
			viper.GetInt("api.port"),
			routes.Routes,
		); err != nil {
			log.Fatalf("api server start error: %v", err)
		}
	}()

	defer func() {
		if err := stg.Close(); err != nil {
			log.Error(err)
		}

		brk.Close()

		if viper.GetBool("clean") {
			if err := os.RemoveAll(viper.GetString("db.file")); err != nil {
				log.Error(err)
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
