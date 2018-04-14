package mail

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, status int) (string, func(t *testing.T)) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))

	return ts.URL, func(t *testing.T) {
		ts.Close()
	}
}

func TestBadRequst(t *testing.T) {
	url, teardown := setup(t, http.StatusBadRequest)
	defer teardown(t)

	viper.Set("mail.endpoint", url)

	err := send("", "", "")
	assert.Equal(t, errors.SendMailError, err)
}

func TestBadEndpoint(t *testing.T) {
	viper.Set("mail.endpoint", "invalid endpoint")

	err := send("", "", "")
	assert.Error(t, err)
}

func TestBadTemplatePath(t *testing.T) {
	err := execute("", "", "", "")
	assert.Error(t, err)
}
