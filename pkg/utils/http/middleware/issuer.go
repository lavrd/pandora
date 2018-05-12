package middleware

import (
	"net/http"
)

// Issuer check that user have right to issue certificate
func Issuer(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var (
		// 	acc = r.Context().Value("acc").(*pb.Account)
		// )
		//
		// if acc.Meta.Type != pb.TypeIssuer {
		// 	errors.Forbidden().Http(w)
		// 	return
		// }
		//
		// h.ServeHTTP(w, r)
	}
}
