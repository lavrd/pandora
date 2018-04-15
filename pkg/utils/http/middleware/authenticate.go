package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/spacelavr/pandora/pkg/api/env"
	"github.com/spacelavr/pandora/pkg/distribution"
	"github.com/spacelavr/pandora/pkg/utils/crypto/jwt"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

// Authenticate authenticate middleware
func Authenticate(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			token string
		)

		w.Header().Add("Content-Type", "application/json")

		if r.Header.Get("Authorization") != "" {
			// parse authorization header
			auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

			// Check authorization header parts length and authorization header format
			if len(auth) != 2 || auth[0] != "Bearer" {
				errors.Unauthorized().Http(w)
				return
			}

			token = auth[1]
		} else {
			errors.Unauthorized().Http(w)
			return
		}

		// validate token
		email, err := jwt.Validate(token)
		if err != nil || email == "" {
			errors.Unauthorized().Http(w)
			return
		}

		dist := distribution.Distribution{env.GetStorage()}

		// fetch account
		acc, err := dist.AccountFetch(email)
		if err != nil {
			errors.InternalServerError().Http(w)
			return
		}
		// if account not in system -> unauthorized
		if acc == nil {
			errors.Unauthorized().Http(w)
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "acc", acc)))
	}
}
