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

func TestResponses(t *testing.T) {
	cases := []struct {
		response *errors.Response
		name     string
		code     int
	}{
		{
			errors.NotImplemented(),
			"NotImplemented",
			http.StatusNotImplemented,
		},
		{
			errors.Unauthorized(),
			"Unauthorized",
			http.StatusUnauthorized,
		},
		{
			errors.InvalidJSON(),
			"InvalidJSON",
			http.StatusUnprocessableEntity,
		},
		{
			errors.NotFound(""),
			"NotFound",
			http.StatusNotFound,
		},
		{
			errors.AlreadyExists(""),
			"AlreadyExists",
			http.StatusConflict,
		},
		{
			errors.InternalServerError(),
			"InternalServerError",
			http.StatusInternalServerError,
		},
		{
			errors.BadParameter(""),
			"BadParameter",
			http.StatusUnprocessableEntity,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			url, teardown := setup(t, c.response)
			defer teardown(t)

			r, err := http.Get(url)
			assert.NoError(t, err)
			defer r.Body.Close()

			assert.Equal(t, c.code, r.StatusCode)
		})
	}
}
