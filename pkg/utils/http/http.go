package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spacelavr/pandora/pkg/utils/log"
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
func Listen(endpoint string, routes []Route) error {
	log.Debugf("listen http server on %s", endpoint)

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./dashboard/static/"))))

	for _, route := range routes {
		r.Handle(route.Path, Handle(route.Handler, route.Middleware...)).Methods(route.Method)
	}

	srv := &http.Server{
		Handler:           r,
		Addr:              endpoint,
		ReadHeaderTimeout: time.Second * 5,
		IdleTimeout:       time.Second * 5,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 5,
	}

	// todo need to regenerate cert and key for other domain
	// todo to config
	return srv.ListenAndServeTLS("./contrib/cert.pem", "./contrib/key.pem")
}

// DefaultHeaders add default headers
func DefaultHeaders(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-type", "application/json")
}
