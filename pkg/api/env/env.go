package env

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/types"
)

var (
	e = &env{}
)

type env struct {
	storage *storage.Storage
	broker  *broker.Broker
	chCert  chan *types.Certificate
}

func init() {
	e.chCert = make(chan *types.Certificate)
}

// SendCert send certificate to cert channel
func SendCert(cert *types.Certificate) {
	e.chCert <- cert
}

// ReadCert read certificate from cert channel
func ReadCert() chan *types.Certificate {
	return e.chCert
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
