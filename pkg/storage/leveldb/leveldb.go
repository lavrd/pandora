package leveldb

import (
	"os"

	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	filepath string
	db       *leveldb.DB
}

func New(filepath string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &LevelDB{
		db:       db,
		filepath: filepath,
	}, nil
}

func (ldb *LevelDB) Close() {
	if err := ldb.db.Close(); err != nil {
		log.Error(err)
	}
}

func (ldb *LevelDB) Clean() error {
	if err := os.RemoveAll(ldb.filepath); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
