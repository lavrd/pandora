package membership

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/membership/rpc"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func Daemon() bool {
	log.Debug("start membership daemon")

	var (
		sig = make(chan os.Signal)
	)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	stg, err := storage.Connect(
		config.Viper.Membership.Database.Endpoint,
		config.Viper.Membership.Database.Database,
		config.Viper.Membership.Database.User,
		config.Viper.Membership.Database.Password,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := rpc.New().Listen(); err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if config.Viper.Runtime.Clean {
			// todo clean database
		}
	}()

	<-sig
	log.Debug("handle SIGINT and SIGTERM")
	return true
}
