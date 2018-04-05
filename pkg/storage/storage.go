package storage

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/viper"
)

type storage struct {
	*bolt.DB
}

// Open connect to storage
func Open() (*storage, error) {

	opts := &bolt.Options{
		Timeout: 1 * time.Second,
	}

	db, err := bolt.Open(viper.GetString("db.file"), 0600, opts)
	if err != nil {
		return nil, err
	}

	return &storage{db}, err
}

// CreateBucket create bucket
func (s *storage) CreateBucket(bucket string) error {
	return s.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		if err != nil {
			log.Error(err)
		}
		return err
	})
}

// Delete bucket delete bucket
func (s *storage) DeleteBucket(bucket string) error {
	return s.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucket))
		if err != nil {
			log.Error(err)
		}
		return err
	})
}

// Delete delete key from bucket
func (s *storage) Delete(bucket, key string) error {
	return s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Delete([]byte(key))
		if err != nil {
			log.Error(err)
		}
		return err
	})
}

// Get returns value by bucket and key
func (s *storage) Get(bucket, key string, data *interface{}) error {
	return s.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		err := json.Unmarshal(v, data)
		if err != nil {
			log.Error(err)
		}
		return err
	})
}

// Put put value in storage by bucket and key
func (s *storage) Put(bucket, key string, value []byte) error {
	return s.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), value)
		if err != nil {
			log.Error(err)
		}
		return err
	})
}
