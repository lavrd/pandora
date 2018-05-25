package arangodb

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

// Storage
type ArangoDB struct {
	database string
	client   driver.Client
}

// Connect connect to database
func Connect(endpoint, database, user, password string) (*ArangoDB, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{endpoint},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(user, password),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	storage := &ArangoDB{database: database, client: client}

	if err = storage.Init(); err != nil {
		return nil, err
	}

	return storage, nil
}

// Init initialize storage
func (s *ArangoDB) Init() error {
	var (
		db driver.Database
	)

	if err := s.InitDatabase(&db); err != nil {
		return err
	}

	if err := s.InitCollections(&db); err != nil {
		return err
	}

	return nil
}

// InitDatabase init storage database
func (s *ArangoDB) InitDatabase(db *driver.Database) error {
	ctx := context.Background()

	ok, err := s.client.DatabaseExists(ctx, s.database)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		*db, err = s.client.CreateDatabase(ctx, s.database, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	} else {
		*db, err = s.Database()
		if err != nil {
			return err
		}
	}

	return nil
}

// InitCollections init collections
func (s *ArangoDB) InitCollections(db *driver.Database) error {
	ctx := context.Background()

	ok, err := (*db).CollectionExists(ctx, CAccount)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		_, err = (*db).CreateCollection(ctx, CAccount, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	ok, err = (*db).CollectionExists(ctx, CCertificates)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		_, err := (*db).CreateCollection(ctx, CCertificates, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

// Clean clean database
func (s *ArangoDB) Clean() error {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return err
	}

	err = db.Remove(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Database returns database
func (s *ArangoDB) Database() (driver.Database, error) {
	ctx := context.Background()

	db, err := s.client.Database(ctx, s.database)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

// Collection returns collection
func (s *ArangoDB) Collection(name string) (driver.Collection, error) {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return nil, err
	}

	col, err := db.Collection(ctx, name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return col, nil
}

// Exec exec query and returns document meta
func (s *ArangoDB) Exec(query string, vars map[string]interface{}, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return nil, err
	}

	cursor, err := db.Query(driver.WithQueryCount(ctx), query, vars)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close()

	if cursor.Count() == 0 {
		return nil, errors.NotFound
	}

	for {
		meta, err := cursor.ReadDocument(ctx, document)
		if err != nil {
			if driver.IsNoMoreDocuments(err) {
				return &meta, nil
			}
			log.Error(err)
			return nil, err
		}
	}
}

func (s *ArangoDB) Read(key, collection string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	db, err := s.Database()
	if err != nil {
		return nil, err
	}

	col, err := db.Collection(ctx, collection)
	if err != nil {
		return nil, nil
	}

	meta, err := col.ReadDocument(ctx, key, document)
	if err != nil {
		if driver.IsNotFound(err) {
			return nil, errors.NotFound
		}

		return nil, err
	}

	return &meta, nil
}

// Write write document to collection
func (s *ArangoDB) Write(collection string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	col, err := s.Collection(collection)
	if err != nil {
		return nil, err
	}

	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &meta, nil
}

// Update update document by collection and key
func (s *ArangoDB) Update(collection, key string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	col, err := s.Collection(collection)
	if err != nil {
		return nil, err
	}

	meta, err := col.UpdateDocument(ctx, key, document)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &meta, nil
}
