package env

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/master/events"
	"github.com/spacelavr/pandora/pkg/master/runtime"
)

var (
	e = &env{}
)

type env struct {
	broker  *broker.Broker
	runtime *runtime.Runtime
	events  *events.Events
}

// SetRuntime set runtime to env
func SetRuntime(rt *runtime.Runtime) {
	e.runtime = rt
}

func SetEvents(events *events.Events) {
	e.events = events
}

func GetEvents() *events.Events {
	return e.events
}

// GetRuntime get runtime from env
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
