package storage

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/viper"
)

const (
	BucketAccount = "account"
)

var (
	db = &bolt.DB{}
)

// Open open database
func Open() (err error) {
	var (
		opts = &bolt.Options{
			Timeout: 1 * time.Second,
		}
	)

	if db, err = bolt.Open(viper.GetString("db.file"), 0600, opts); err != nil {
		log.Error(err)
		return
	}

	if err = initBuckets(); err != nil {
		return
	}

	return
}

func initBuckets() error {
	if err := createBucket(BucketAccount); err != nil {
		return err
	}
	return nil
}

// Close close database
func Close() error {
	return db.Close()
}

func createBucket(bucket string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
	if err != nil {
		log.Error(err)
	}
	return err
}

func delete(bucket, key string) error {
	err := db.Update(func(tx *bolt.Tx) error {
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

func get(bucket, key string, data interface{}) error {
	err := db.View(func(tx *bolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			if v := b.Get([]byte(key)); v != nil {
				return json.Unmarshal(v, data)
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

func put(bucket, key string, value interface{}) error {
	if data, err := json.Marshal(value); err == nil {
		err = db.Update(func(tx *bolt.Tx) error {
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
