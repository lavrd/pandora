package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/spacelavr/pandora/pkg/api/routes"
	"github.com/stretchr/testify/assert"
)

func TestHealthH(t *testing.T) {

	t.Parallel()

	cases := []struct {
		handler func(http.ResponseWriter, *http.Request)
		code    int
		name    string
	}{
		{
			name:    "checking health handler",
			handler: routes.HealthH,
			code:    http.StatusNotImplemented,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/health", nil)
			assert.NoError(t, err)

			router := mux.NewRouter()
			router.HandleFunc("/health", c.handler)

			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			assert.Equal(t, c.code, res.Code)
		})
	}
}
