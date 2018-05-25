package node

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/runtime"
	"github.com/spacelavr/pandora/pkg/storage"
)

var (
	e = &env{}
)

type env struct {
	broker  *broker.Broker
	storage *storage.Storage
	runtime *runtime.Runtime
}

// SetStorage set storage to env
func SetStorage(stg *storage.Storage) {
	e.storage = stg
}

// GetStorage get storage from env
func GetStorage() *storage.Storage {
	return e.storage
}

// SetStorage set storage to env
func SetRuntime(r *runtime.Runtime) {
	e.runtime = r
}

// GetStorage get storage from env
func GetRuntime() *runtime.Runtime {
	return e.runtime
}

// SetBroker set broker to env
func SetBroker(brk *broker.Broker) {
	e.broker = brk
}

// GetBroker returns broker from env
func GetBroker() *broker.Broker {
	return e.broker
}

