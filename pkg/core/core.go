package core

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/log"
)

// Daemon start core daemon
func Daemon() bool {
	log.Debug("start core daemon")

	var (
		sig = make(chan os.Signal, 1)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
