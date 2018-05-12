package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func test() {
	leveldb.OpenFile("", nil)
}
