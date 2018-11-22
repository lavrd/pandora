package membership

import (
	"os"
	"os/signal"
	"syscall"

	"pandora/pkg/conf"
	"pandora/pkg/membership/env"
	"pandora/pkg/membership/rpc"
	"pandora/pkg/storage/arangodb"
	"pandora/pkg/utils/log"
)

const (
	MEMBERSHIP = "membership"
)

// Daemon start daemon service
func Daemon() bool {
	log.Debug("start membership daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	stg, err := arangodb.New(
		conf.Conf.Arangodb.Endpoint,
		conf.Conf.Arangodb.Database,
		conf.Conf.Arangodb.User,
		conf.Conf.Arangodb.Password,
	)
	if err != nil {
		log.Fatal(err)
	}

	r, err := rpc.New()
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	env.SetStorage(stg)

	go func() {
		if err := r.Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if conf.Conf.Runtime.Clean {
			if err := stg.Clean(); err != nil {
				log.Error(err)
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
