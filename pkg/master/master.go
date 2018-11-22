package master

import (
	"os"
	"os/signal"
	"syscall"

	"pandora/pkg/blockchain"
	"pandora/pkg/broker"
	"pandora/pkg/master/env"
	"pandora/pkg/master/events"
	"pandora/pkg/master/rpc"
	"pandora/pkg/utils/log"
)

const (
	MASTER = "master"
)

// Daemon start master daemon
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
