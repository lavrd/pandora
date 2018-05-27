package env

import (
	"github.com/spacelavr/pandora/pkg/storage/arangodb"
)

var (
	e = &env{}
)

type env struct {
	storage *arangodb.ArangoDB
}

func SetStorage(stg *arangodb.ArangoDB) {
	e.storage = stg
}

func GetStorage() *arangodb.ArangoDB {
	return e.storage
}
