package discovery

import (
	"os"
	"os/signal"
	"syscall"

	"pandora/pkg/discovery/rpc"
	"pandora/pkg/utils/log"
)

const (
	DISCOVERY = "discovery"
)

// Daemon start discovery daemon
func Daemon() bool {
	log.Debug("start discovery daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := rpc.New().Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
