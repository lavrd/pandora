package storage

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/spacelavr/pandora/pkg/log"
)

// Storage
type Storage struct {
	database string
	client   driver.Client
}

// ConnectOpts
type ConnectOpts struct {
	Endpoint string
	User     string
	Password string
	Database string
}

// Connect connect to database
func Connect(opts *ConnectOpts) (*Storage, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{opts.Endpoint},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(opts.User, opts.Password),
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	storage := &Storage{database: opts.Database, client: client}

	if err = storage.Init(); err != nil {
		return nil, err
	}

	return storage, nil
}

// Close close connection with database
func (s *Storage) Close() error {
	return nil
}

// Init initialize database
func (s *Storage) Init() error {
	var (
		ctx = context.Background()
		db  driver.Database
	)

	ok, err := s.client.DatabaseExists(ctx, s.database)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		db, err = s.client.CreateDatabase(ctx, s.database, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	} else {
		db, err = s.Database()
		if err != nil {
			return err
		}
	}

	ok, err = db.CollectionExists(ctx, CollectionAccount)
	if err != nil {
		log.Error(err)
		return err
	}
	if !ok {
		_, err = db.CreateCollection(ctx, CollectionAccount, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

// Clean clean database
func (s *Storage) Clean() error {
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
func (s *Storage) Database() (driver.Database, error) {
	ctx := context.Background()

	db, err := s.client.Database(ctx, s.database)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

// Collection returns collection
func (s *Storage) Collection(name string) (driver.Collection, error) {
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

// Read read document by collection and key
func (s *Storage) Read(collection, key string, document interface{}) (*driver.DocumentMeta, error) {
	ctx := context.Background()

	col, err := s.Collection(collection)
	if err != nil {
		return nil, err
	}

	meta, err := col.ReadDocument(ctx, key, document)
	if err != nil {
		if !driver.IsNotFound(err) {
			log.Error(err)
		}
		return nil, err
	}

	return &meta, nil
}

// Write write document to collection
func (s *Storage) Write(collection string, document interface{}) (*driver.DocumentMeta, error) {
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
func (s *Storage) Update(collection, key string, document interface{}) (*driver.DocumentMeta, error) {
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
