package env

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/storage"
)

var (
	e = &env{}
)

type env struct {
	storage *storage.Storage
	broker  *broker.Broker
}

// SetBroker set broker to env
func SetBroker(brk *broker.Broker) {
	e.broker = brk
}

// GetBroker returns broker from env
func GetBroker() *broker.Broker {
	return e.broker
}

// SetStorage set storage to env
func SetStorage(stg *storage.Storage) {
	e.storage = stg
}

// GetStorage returns storage from env
func GetStorage() *storage.Storage {
	return e.storage
}
