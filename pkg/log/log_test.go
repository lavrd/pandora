package log_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetVerbose(true)
}

func setup(t *testing.T) (*os.File, *os.File, func(t *testing.T)) {
	stdout := os.Stdout
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdout = w

	return r, w, func(t *testing.T) {
		os.Stdout = stdout
		r.Close()
		w.Close()
	}
}

func read(r *os.File) string {
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	return <-outC
}

func find(out, pattern string, offset int) string {
	start := strings.Index(out, pattern)
	stop := start + offset
	return out[start:stop]
}

func TestHttp(t *testing.T) {
	r, w, teardown := setup(t)
	defer teardown(t)

	var (
		start     = time.Now()
		reqTime   = start.UTC().Format(time.RFC1123Z)
		duration  = time.Since(start).String()
		ip        = "127.0.0.1"
		method    = "GET"
		route     = "/health"
		proto     = "HTTP/1.1"
		userAgent = "curl"
		code      = 501
		size      = 40
	)

	log.Http(w, ip, reqTime, method, route, proto, duration, userAgent, code, size)

	w.Close()
	out := read(r)

	assert.Equal(
		t,
		fmt.Sprintf(log.CommonLogFormat, ip, "-", "-", reqTime, method, route, proto, code, size, duration, userAgent),
		out,
	)
}

func TestDebug(t *testing.T) {
	r, w, teardown := setup(t)
	defer teardown(t)

	log.SetOut(w)
	log.Debug("debug log")

	w.Close()
	out := read(r)

	assert.Equal(t, "level=debug msg=\"debug log\"", find(out, "level", 27))
}

func TestError(t *testing.T) {
	r, w, teardown := setup(t)
	defer teardown(t)

	log.SetOut(w)
	log.Error("error log")

	w.Close()
	out := read(r)

	assert.Equal(t, "level=error msg=\"error log\"", find(out, "level", 27))
}

func TestErrorf(t *testing.T) {
	r, w, teardown := setup(t)
	defer teardown(t)

	log.SetOut(w)
	log.Errorf("%s", "formatted error log")

	w.Close()
	out := read(r)

	assert.Equal(t, "level=error msg=\"formatted error log\"", find(out, "level", 37))
}

func TestDebugf(t *testing.T) {
	r, w, teardown := setup(t)
	defer teardown(t)

	log.SetOut(w)
	log.Debugf("%s", "formatted debug log")

	w.Close()
	out := read(r)

	assert.Equal(t, "level=debug msg=\"formatted debug log\"", find(out, "level", 37))
}
