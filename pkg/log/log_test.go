//
// do not run these tests in parallel
// because threads of stdout intersect
// and the data is distorted
//

package log_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.Init(true)
}

func TestDebug(t *testing.T) {

	r, w, teardown := setup(t)
	defer teardown(t)

	log.SetOut(w)
	log.Debug("debug log")

	w.Close()
	out := read(r)

	assert.Equal(t, "level=debug msg=\"[debug log]\"", out[33:62])
}

func TestError(t *testing.T) {

	r, w, teardown := setup(t)
	defer teardown(t)

	log.SetOut(w)
	log.Error("error log")

	w.Close()
	out := read(r)

	assert.Equal(t, "level=error msg=\"[error log]\"", out[33:62])
}

func setup(t *testing.T) (*os.File, *os.File, func(t *testing.T)) {

	stdout := os.Stdout
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdout = w

	return r, w, func(t *testing.T) {
		os.Stdout = stdout
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
