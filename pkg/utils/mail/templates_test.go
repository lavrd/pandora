package mail_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/mail"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T, subject string) (string, func(t *testing.T)) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := &mail.Mail{}

		err := json.NewDecoder(r.Body).Decode(m)
		assert.NoError(t, err)
		assert.Equal(t, subject, m.Subject)

		w.WriteHeader(http.StatusAccepted)
	}))

	return ts.URL, func(t *testing.T) {
		ts.Close()
	}
}

func TestSendAccountCreated(t *testing.T) {
	url, teardown := setup(t, viper.GetString("mail.subjects.account.created"))
	defer teardown(t)

	viper.Set("mail.templates.account.created", "./templates/account/created.html")
	viper.Set("mail.endpoint", url)

	err := mail.SendAccountCreated("", "")
	assert.NoError(t, err)
}

func TestSendAccountRecovery(t *testing.T) {
	url, teardown := setup(t, viper.GetString("mail.subjects.account.recovery"))
	defer teardown(t)

	viper.Set("mail.templates.account.recovery", "./templates/account/recovery.html")
	viper.Set("mail.endpoint", url)

	err := mail.SendAccountRecovery("", "")
	assert.NoError(t, err)
}
