package env

import (
	"github.com/spacelavr/pandora/pkg/blockchain"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage/leveldb"
)

var (
	e = &env{}
)

type env struct {
	storage *leveldb.LevelDB
	rpc     *rpc.RPC
	key     *pb.PublicKey
	bc      *blockchain.Blockchain
}

func SetBlockchain(bc *blockchain.Blockchain) {
	e.bc = bc
}

func GetBlockchain() *blockchain.Blockchain {
	return e.bc
}

func SetStorage(stg *leveldb.LevelDB) {
	e.storage = stg
}

func GetStorage() *leveldb.LevelDB {
	return e.storage
}

func SetRPC(rpc *rpc.RPC) {
	e.rpc = rpc
}

func GetRPC() *rpc.RPC {
	return e.rpc
}

func SetKey(key *pb.PublicKey) {
	e.key = key
}

func GetKey() *pb.PublicKey {
	return e.key
}
