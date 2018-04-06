package storage_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	bucket = "bucket"
	prefix = "prefix"
	dir    = "./"
	key    = "key"
)

func setup(t *testing.T) (*storage.Store, func(t *testing.T)) {

	file, err := ioutil.TempFile(dir, prefix)
	assert.NoError(t, err)
	file.Close()

	viper.Set("db.file", file.Name())

	store, err := storage.Open()
	assert.NoError(t, err)
	assert.NotNil(t, store.DB)

	return store, func(t *testing.T) {
		store.Close()
		os.Remove(file.Name())
	}
}

func TestOpen(t *testing.T) {

	t.Parallel()

	_, teardowm := setup(t)
	defer teardowm(t)
}

func TestStorage_CreateBucket(t *testing.T) {

	t.Parallel()

	store, teardown := setup(t)
	defer teardown(t)

	err := store.CreateBucket(bucket)
	assert.NoError(t, err)

	err = store.CreateBucket(bucket)
	assert.Error(t, err)
}

func TestStorage_DeleteBucket(t *testing.T) {

	t.Parallel()

	store, teardown := setup(t)
	defer teardown(t)

	err := store.DeleteBucket(bucket)
	assert.Error(t, err)

	err = store.CreateBucket(bucket)
	assert.NoError(t, err)

	err = store.DeleteBucket(bucket)
	assert.NoError(t, err)
}

func TestStorage_Put(t *testing.T) {

	t.Parallel()

	store, teardown := setup(t)
	defer teardown(t)

	err := store.Put(bucket, key, "")
	assert.NoError(t, err)

	err = store.CreateBucket(bucket)
	assert.NoError(t, err)

	data := &struct {
		data string
	}{
		data: "test",
	}

	err = store.Put(bucket, key, data)
	assert.NoError(t, err)

	err = store.Put(bucket, key, data)
	assert.NoError(t, err)
}

func TestStorage_Get(t *testing.T) {

	t.Parallel()

	store, teardown := setup(t)
	defer teardown(t)

	err := store.CreateBucket(bucket)
	assert.NoError(t, err)

	type TestStruct struct {
		Data string `json:"data"`
	}

	data := new(TestStruct)

	err = store.Get(bucket, key, data)
	assert.NoError(t, err)
	assert.Equal(t, "", data.Data)

	data.Data = "test"

	err = store.Put(bucket, key, data)
	assert.NoError(t, err)

	data = new(TestStruct)

	err = store.Get(bucket, key, data)
	assert.NoError(t, err)
	assert.Equal(t, "test", data.Data)

	err = store.Get("undefined", "", "")
	assert.NoError(t, err)
}

func TestStorage_Delete(t *testing.T) {

	t.Parallel()

	store, teardown := setup(t)
	defer teardown(t)

	err := store.Delete(bucket, key)
	assert.NoError(t, err)

	err = store.CreateBucket(bucket)
	assert.NoError(t, err)

	err = store.Put(bucket, key, "")
	assert.NoError(t, err)

	err = store.Delete(bucket, key)
	assert.NoError(t, err)
}
