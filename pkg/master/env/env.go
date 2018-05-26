package env

import (
	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/master/events"
)

var (
	e = &env{}
)

type env struct {
	bc  *blockchain.Blockchain
	evt *events.Events
}

func SetBlockchain(bc *blockchain.Blockchain) {
	e.bc = bc
}

func SetEvents(evt *events.Events) {
	e.evt = evt
}

func GetEvents() *events.Events {
	return e.evt
}

func GetBlockchain() *blockchain.Blockchain {
	return e.bc
}
