package env

import (
	"pandora/pkg/storage/arangodb"
)

var (
	e = &env{}
)

type env struct {
	storage *arangodb.Arangodb
}

// SetStorage set storage to env
func SetStorage(stg *arangodb.Arangodb) {
	e.storage = stg
}

// GetStorage get storage from env
func GetStorage() *arangodb.Arangodb {
	return e.storage
}
