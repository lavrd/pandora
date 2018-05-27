package env

import (
	"github.com/spacelavr/pandora/pkg/membership/rpc"
	"github.com/spacelavr/pandora/pkg/storage/arangodb"
)

var (
	e = &env{}
)

type env struct {
	storage *arangodb.ArangoDB
	rpc     *rpc.RPC
}

func SetStorage(stg *arangodb.ArangoDB) {
	e.storage = stg
}

func GetStorage() *arangodb.ArangoDB {
	return e.storage
}

func SetRPC(rpc *rpc.RPC) {
	e.rpc = rpc
}

func GetRPC() *rpc.RPC {
	return e.rpc
}
