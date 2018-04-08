package errors_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, r *errors.Response) (string, func(t *testing.T)) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		r.Http(w)
	}))

	return ts.URL, func(t *testing.T) {
		ts.Close()
	}
}

func TestNotImplemented(t *testing.T) {

	t.Parallel()

	url, teardown := setup(t, errors.NotImplemented())
	defer teardown(t)

	r, err := http.Get(url)
	assert.NoError(t, err)
	defer r.Body.Close()

	assert.Equal(t, http.StatusNotImplemented, r.StatusCode)
}
