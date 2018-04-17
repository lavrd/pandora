package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/utils/http/middleware"
)

// Route
type Route struct {
	Path       string
	Handler    func(w http.ResponseWriter, r *http.Request)
	Middleware []Middleware
	Method     string
}

// Middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

var (
	GET  = http.MethodGet
	POST = http.MethodPost
)

// Handle prepare handler
func Handle(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	headers := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			DefaultHeaders(w, r)
			h.ServeHTTP(w, r)
		}
	}

	h = headers(h)
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

// Listen start listen http requests
func Listen(port int, routes []Route) error {
	log.Debugf("listen http server on :%d", port)

	router := mux.NewRouter()

	for _, route := range routes {
		router.Handle(route.Path, Handle(route.Handler, route.Middleware...)).Methods(route.Method)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", port), middleware.Logger(router))
}

// DefaultHeaders add default headers
func DefaultHeaders(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-type", "application/json")
}
