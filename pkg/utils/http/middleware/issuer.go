package middleware

import (
	"context"
	"net/http"

	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

// Issuer check that user have right to issue certificate
func Issuer(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			acc = r.Context().Value("acc").(types.Account)
		)

		if acc.Type != types.TypeIssuer {
			errors.Forbidden().Http(w)
			return
		}

		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "acc", acc)))
	}
}
