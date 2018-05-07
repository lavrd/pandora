package broker

import (
	"github.com/nats-io/gnatsd/server"
	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func Daemon() error {
	s := server.New(&server.Options{
		Port: config.Viper.Validator.Broker.Port,
	})
	if s == nil {
		log.Error(errors.UnableStartGNATSD)
		return errors.UnableStartGNATSD
	}

	s.Start()
	defer s.Shutdown()

	return nil
}
