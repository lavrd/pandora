package env

import (
	"github.com/spacelavr/pandora/pkg/broker"
)

var (
	e = &env{}
)

type env struct {
	broker *broker.Broker
}

// SetBroker set broker to env
func SetBroker(brk *broker.Broker) {
	e.broker = brk
}

// GetBroker returns broker from env
func GetBroker() *broker.Broker {
	return e.broker
}
