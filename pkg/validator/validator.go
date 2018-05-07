package validator

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/validator/broker"
	"github.com/spacelavr/pandora/pkg/validator/rpc"
)

func Daemon() bool {
	log.Debug("start validator daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := broker.Daemon(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := rpc.NotifyTracker(); err != nil {
		log.Fatal(err)
	}

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
