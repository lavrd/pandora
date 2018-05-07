package node

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start node daemon
func Daemon() bool {
	log.Debug("start node daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	brkOpts, err := rpc.GetValidators()
	if err != nil {
		log.Fatal(err)
	}

	brk, err := broker.Connect(&broker.Opts{
		Endpoint: brkOpts.Endpoint,
		User:     brkOpts.User,
		Password: brkOpts.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	env.SetBroker(brk)

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for t := range time.Tick(time.Second * 1) {
			fmt.Println(1)
			events.ChSendNewBlock <- &types.Block{
				Index: t.Second(),
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
