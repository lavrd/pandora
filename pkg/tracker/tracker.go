package tracker

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/tracker/rpc"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func Daemon() bool {
	log.Debug("start tracker daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := rpc.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
