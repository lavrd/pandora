package node

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/routes"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/node/runtime"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Daemon start node daemon
func Daemon() bool {
	log.Debug("start node daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	candidate := request.Candidate{
		Name:  &config.Viper.Node.Meta.Name,
		Email: &config.Viper.Node.Meta.Email,
	}
	if err := candidate.Validate(); err != nil {
		log.Fatal(err.Message)
	}

	log.Debug(candidate)
	// todo with either start or once?
	key, err := rpc.NodeReg(&pb.Candidate{
		Email: *candidate.Email,
		Name:  *candidate.Name,
	})
	if err != nil {
		if err != errors.AlreadyExists {
			log.Fatal(err)
		}
	}

	log.Debug(key)

	netOpts, err := rpc.Network()
	if err != nil {
		log.Fatal(err)
	}

	brk, err := broker.Connect(&broker.Opts{
		Endpoint: netOpts.Broker.Endpoint,
		User:     netOpts.Broker.User,
		Password: netOpts.Broker.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer brk.Close()

	rt, err := runtime.New(key)
	if err != nil {
		log.Fatal(err)
	}

	rt.PublicKey = key.PublicKey
	rt.FullName = config.Viper.Node.Meta.Name

	env.SetBroker(brk)
	env.SetRuntime(rt)

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := http.Listen(fmt.Sprintf(":%d", config.Viper.Node.Port), routes.Routes); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if config.Viper.Runtime.Clean {
			if err := os.RemoveAll(config.Viper.Node.Database.FilePath); err != nil {
				log.Error(err)
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
