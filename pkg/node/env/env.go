package env

import (
	"pandora/pkg/blockchain"
	"pandora/pkg/node/rpc"
	"pandora/pkg/pb"
	"pandora/pkg/storage/leveldb"
)

var (
	e = &env{}
)

type env struct {
	storage *leveldb.Leveldb
	rpc     *rpc.RPC
	key     *pb.PublicKey
	bc      *blockchain.Blockchain
}

// SetBlockchain set blockchain to env
func SetBlockchain(bc *blockchain.Blockchain) {
	e.bc = bc
}

// GetBlockchain get blockchain from env
func GetBlockchain() *blockchain.Blockchain {
	return e.bc
}

// SetStorage set storage to env
func SetStorage(stg *leveldb.Leveldb) {
	e.storage = stg
}

// GetStorage get storage from env
func GetStorage() *leveldb.Leveldb {
	return e.storage
}

// SetRPC set rpc to env
func SetRPC(rpc *rpc.RPC) {
	e.rpc = rpc
}

// GetRPC get rpc from env
func GetRPC() *rpc.RPC {
	return e.rpc
}

// SetKey set key to env
func SetKey(key *pb.PublicKey) {
	e.key = key
}

// GetKey get key from env
func GetKey() *pb.PublicKey {
	return e.key
}
