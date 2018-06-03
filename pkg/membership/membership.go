package membership

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/membership/env"
	"github.com/spacelavr/pandora/pkg/membership/rpc"
	"github.com/spacelavr/pandora/pkg/storage/arangodb"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

const (
	MEMBERSHIP = "membership"
)

func Daemon() bool {
	log.Debug("start membership daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	stg, err := arangodb.New(
		conf.Viper.Membership.Database.Endpoint,
		conf.Viper.Membership.Database.Database,
		conf.Viper.Membership.Database.User,
		conf.Viper.Membership.Database.Password,
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
		if conf.Viper.Runtime.Clean {
			if err := stg.Clean(); err != nil {
				log.Error(err)
			}
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
