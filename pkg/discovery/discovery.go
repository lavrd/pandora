package discovery

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/discovery/env"
	"github.com/spacelavr/pandora/pkg/discovery/http/routes"
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start discovery daemon
func Daemon() bool {
	log.Debug("start discovery daemon")

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

	env.SetBroker(brk)

	go func() {
		if err := http.Listen(config.Viper.Discovery.Port, routes.Routes); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
