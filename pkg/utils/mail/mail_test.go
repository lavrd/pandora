package mail

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/stretchr/testify/assert"
)

func setup(_ *testing.T, status int) (string, func(t *testing.T)) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
	}))

	return ts.URL, func(t *testing.T) {
		ts.Close()
	}
}

func TestBadRequest(t *testing.T) {
	url, teardown := setup(t, http.StatusBadRequest)
	defer teardown(t)

	conf.Viper.Membership.Mail.Send = true
	conf.Viper.Membership.Mail.Endpoint = url

	err := send("", "", "")
	assert.Equal(t, errors.SendMailError, err)
}

func TestBadEndpoint(t *testing.T) {
	conf.Viper.Membership.Mail.Send = true
	conf.Viper.Membership.Mail.Endpoint = "invalid endpoint"

	err := send("", "", "")
	assert.Error(t, err)
}

func TestBadTemplatePath(t *testing.T) {
	err := execute("", "", "", "")
	assert.Error(t, err)
}
