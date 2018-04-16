package storage

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spacelavr/pandora/pkg/log"
)

// Storage
type Storage struct {
	*bolt.DB
}

// Open open database
func Open(path string) (*Storage, error) {
	var (
		opts = &bolt.Options{
			Timeout: 1 * time.Second,
		}
	)

	db, err := bolt.Open(path, 0600, opts)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	s := &Storage{db}

	if err = s.InitBuckets(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) InitBuckets() error {
	if err := s.CreateBucket(BucketAccount); err != nil {
		return err
	}
	return nil
}

// Close close database
func (s *Storage) Close() error {
	return s.DB.Close()
}

// CreateBucket create bucket
func (s *Storage) CreateBucket(bucket string) error {
	err := s.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	if err != nil {
		log.Error(err)
	}
	return err
}

// Delete delete key from bucket
func (s *Storage) Delete(bucket, key string) error {
	err := s.Update(func(tx *bolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			return b.Delete([]byte(key))
		}
		return nil
	})
	if err != nil {
		log.Error(err)
	}
	return err
}

// Get returns value by key and bucket
func (s *Storage) Get(bucket, key string, value interface{}) error {
	err := s.View(func(tx *bolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			if v := b.Get([]byte(key)); v != nil {
				return json.Unmarshal(v, value)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		log.Error(err)
	}
	return err
}

// Put put value by key and bucket
func (s *Storage) Put(bucket, key string, value interface{}) error {
	if data, err := json.Marshal(value); err == nil {
		err = s.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			err := b.Put([]byte(key), data)
			return err
		})
		if err != nil {
			log.Error(err)
		}
		return err
	} else {
		log.Error(err)
		return err
	}
}
