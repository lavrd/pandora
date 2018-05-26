package leveldb

import (
	"os"

	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	path string
	db   *leveldb.DB
}

func New(path string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &LevelDB{
		db:   db,
		path: path,
	}, nil
}

func (ldb *LevelDB) Close() error {
	if err := ldb.db.Close(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ldb *LevelDB) Clean() error {
	if err := os.RemoveAll(ldb.path); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
