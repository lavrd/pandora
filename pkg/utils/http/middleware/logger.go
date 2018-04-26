package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/spacelavr/pandora/pkg/utils/log"
)

type loggedResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

// Logger logger middleware
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := newLoggedResponseWriter(w)
		h.ServeHTTP(lrw, r)

		log.Http(
			os.Stdout,
			r.RemoteAddr,
			start.UTC().Format(time.RFC1123Z),
			r.Method,
			r.URL.Path,
			r.Proto,
			time.Since(start).String(),
			r.UserAgent(),
			lrw.status,
			lrw.size,
		)
	})
}

// WriteHeader redefined http write header
func (c *loggedResponseWriter) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

// Write redefined http write
func (c *loggedResponseWriter) Write(b []byte) (int, error) {
	size, err := c.ResponseWriter.Write(b)
	c.size += size
	return size, err
}

func newLoggedResponseWriter(w http.ResponseWriter) *loggedResponseWriter {
	return &loggedResponseWriter{
		ResponseWriter: w,
	}
}
