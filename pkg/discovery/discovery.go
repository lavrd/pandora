package discovery

import (
	"pandora/pkg/discovery/rpc"
	"pandora/pkg/utils/log"
	"pandora/pkg/utils/signalutils"
)

const (
	DISCOVERY = "discovery"
)

// Daemon start discovery daemon
func Daemon() bool {
	log.Debug("start discovery daemon")

	go func() {
		if err := rpc.New().Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	signalutils.NotifyInterrupt()
	return true
}
