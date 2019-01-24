package leveldb

import (
	"os"

	"github.com/syndtr/goleveldb/leveldb"

	"pandora/pkg/utils/errors"
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
	return &Leveldb{filepath, db}, nil
}

// Close close conn with leveldb
func (ldb *Leveldb) Close() error {
	return errors.WithStack(ldb.db.Close())
}

// Clean clean leveldb
func (ldb *Leveldb) Clean() error {
	return errors.WithStack(os.RemoveAll(ldb.filepath))
}
