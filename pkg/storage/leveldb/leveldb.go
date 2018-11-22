package leveldb

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"

	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
)

// Leveldb
type Leveldb struct {
	filepath string
	db       *leveldb.DB
}

// New returns new leveldb
func New(filepath string) (*Leveldb, error) {
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Leveldb{
		db:       db,
		filepath: filepath,
	}, nil
}

// Close close conn with leveldb
func (ldb *Leveldb) Close() {
	if err := ldb.db.Close(); err != nil {
		log.Error(errors.WithStack(err))
	}
}

// Clean clean leveldb
func (ldb *Leveldb) Clean() error {
	if err := os.RemoveAll(ldb.filepath); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
