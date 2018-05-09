package env

import (
	"github.com/spacelavr/pandora/pkg/membership/runtime"
	"github.com/spacelavr/pandora/pkg/storage"
)

var (
	e = &env{}
)

type env struct {
	runtime *runtime.Runtime
	storage *storage.Storage
}

func SetStorage(stg *storage.Storage) {
	e.storage = stg
}

func GetStorage() *storage.Storage {
	return e.storage
}

func SetRuntime(rt *runtime.Runtime) {
	e.runtime = rt
}

func GetRuntime() *runtime.Runtime {
	return e.runtime
}
