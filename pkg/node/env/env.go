package env

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/storage"
)

var (
	e = &env{
		chAccount: make(chan *request.Account),
	}
)

type env struct {
	chAccount chan *request.Account
	broker    *broker.Broker
	storage   *storage.Storage
}

func SendAccount(acc *request.Account) {
	e.chAccount <- acc
}

func ReadAccount() chan *request.Account {
	return e.chAccount
}

// SetStorage set storage to env
func SetStorage(stg *storage.Storage) {
	e.storage = stg
}

// GetStorage get storage from env
func GetStorage() *storage.Storage {
	return e.storage
}

// SetBroker set broker to env
func SetBroker(brk *broker.Broker) {
	e.broker = brk
}

// GetBroker returns broker from env
func GetBroker() *broker.Broker {
	return e.broker
}
