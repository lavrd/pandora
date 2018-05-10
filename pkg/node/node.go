package node

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/routes"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/node/runtime"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage"
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
		FullName: &config.Viper.Node.FullName,
		Email:    &config.Viper.Node.Email,
	}
	if err := candidate.Validate(); err != nil {
		log.Fatal(err.Message)
	}

	// todo with either start or once?
	key, err := rpc.NodeReg(&pb.Candidate{
		Email:    *candidate.Email,
		FullName: *candidate.FullName,
	})
	if err != nil {
		if err != errors.AlreadyExists {
			log.Fatal(err)
		}
	}

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

	stg, err := storage.Connect(&storage.Opts{
		Endpoint: config.Viper.Database.Endpoint,
		Database: config.Viper.Database.Database,
		User:     config.Viper.Database.User,
		Password: config.Viper.Database.Password,
	})
	if err != nil {
		log.Fatal(err)
	}

	rt, err := runtime.New(key)
	if err != nil {
		log.Fatal(err)
	}

	env.SetStorage(stg)
	env.SetBroker(brk)
	env.SetRuntime(rt)

	go func() {
		if err := events.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := http.Listen(config.Viper.Node.Endpoint, routes.Routes); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if config.Viper.Runtime.Clean {
			if err := stg.Clean(); err != nil {
				log.Error(err)
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println(env.GetRuntime().MC)
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
