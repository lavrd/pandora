package mail_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/mail"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {

	t.Parallel()

	cases := []struct {
		status   int
		expected error
		endpoint string
		name     string
	}{
		{
			name:     "error sending mail",
			status:   http.StatusBadRequest,
			expected: errors.New("send mail error"),
		}, {
			name:     "success sending mail",
			status:   http.StatusAccepted,
			expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(c.status)
			}))
			defer ts.Close()

			viper.Set("sendgrid.endpoint", ts.URL)

			err := mail.Send("", "", "")
			assert.Equal(t, c.expected, err)
		})
	}
}
