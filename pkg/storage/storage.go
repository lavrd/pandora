package storage

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/viper"
)

// Store
type Store struct {
	*bolt.DB
}

// Open connect to Store
func Open() (*Store, error) {

	opts := &bolt.Options{
		Timeout: 1 * time.Second,
	}

	db, err := bolt.Open(viper.GetString("db.file"), 0600, opts)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Store{db}, nil
}

// CreateBucket create bucket
func (s *Store) CreateBucket(bucket string) error {
	return s.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		if err != nil {
			log.Error(err)
		}
		return err
	})
}

// DeleteBucket bucket delete bucket
func (s *Store) DeleteBucket(bucket string) error {
	return s.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucket))
		if err != nil {
			log.Error(err)
		}
		return err
	})
}

// Delete delete key from bucket
func (s *Store) Delete(bucket, key string) error {
	return s.Update(func(tx *bolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			err := b.Delete([]byte(key))
			if err != nil {
				log.Error(err)
			}
			return err
		}
		return nil
	})
}

// Get unmarshal value by bucket and key
func (s *Store) Get(bucket, key string, data interface{}) error {
	return s.View(func(tx *bolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			if v := b.Get([]byte(key)); v != nil {
				err := json.Unmarshal(v, data)
				if err != nil {
					log.Error(err)
				}
				return err
			}
			return nil
		}
		return nil
	})
}

// Put put value in Store by bucket and key
func (s *Store) Put(bucket, key string, value interface{}) error {
	if data, err := json.Marshal(value); err == nil {
		return s.Update(func(tx *bolt.Tx) error {
			if b := tx.Bucket([]byte(bucket)); b != nil {
				err := b.Put([]byte(key), data)
				if err != nil {
					log.Error(err)
				}
				return err
			}
			return nil
		})
	} else {
		return err
	}
}
