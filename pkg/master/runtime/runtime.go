package runtime

import (
	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/master/events"
)

var (
	r = &runtime{}
)

type runtime struct {
	Chain  *blockchain.Blockchain
	Events *events.Events
}

func Set(chain *blockchain.Blockchain, events *events.Events) {
	r.Events = events
	r.Chain = chain
}

func Get() *runtime {
	return r
}
