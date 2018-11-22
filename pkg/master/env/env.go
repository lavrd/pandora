package env

import (
	"pandora/pkg/blockchain"
	"pandora/pkg/master/events"
)

var (
	e = &env{}
)

type env struct {
	bc  *blockchain.Blockchain
	evt *events.Events
}

// SetBlockchain set blockchain to env
func SetBlockchain(bc *blockchain.Blockchain) {
	e.bc = bc
}

// SetBlockchain set events to env
func SetEvents(evt *events.Events) {
	e.evt = evt
}

// GetEvents returns events from env
func GetEvents() *events.Events {
	return e.evt
}

// GetBlockchain returns blockchain from env
func GetBlockchain() *blockchain.Blockchain {
	return e.bc
}
