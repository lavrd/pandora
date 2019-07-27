package signalutils

import (
	"os"
	"os/signal"
	"syscall"

	"pandora/pkg/utils/log"
)

func NotifyInterrupt() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-sig
	log.Debug("handle SIGINT, SIGTERM, SIGKILL, SIGQUIT")
}
