package master

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/master/events"
	"github.com/spacelavr/pandora/pkg/master/rpc"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

const (
	Master = "master"
)

func Daemon() bool {
	log.Debug("start master daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	r := rpc.New()

	brkOpts, err := r.InitMaster()
	if err != nil {
		log.Fatal(err)
	}

	brk, err := broker.New(
		brkOpts.Endpoint,
		brkOpts.User,
		brkOpts.Password,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	evt, err := events.New(brk)
	if err != nil {
		log.Fatal(err)
	}

	env.SetBlockchain(blockchain.New())
	env.SetEvents(evt)

	go func() {
		if err := r.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
