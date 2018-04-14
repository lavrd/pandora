package filepath_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/filepath"
	"github.com/stretchr/testify/assert"
)

func TestPKG(t *testing.T) {
	cases := []struct {
		expected string
		path     string
	}{{
		"/utils/filepath/filepath.go",
		"/opt/go/src/github.com/spacelavr/pandora/pkg/utils/filepath/filepath.go",
	}, {
		"/utils/http/http.go",
		"/opt/go/src/github.com/spacelavr/pandora/pkg/utils/http/http.go",
	}, {
		"/client/client.go",
		"/opt/go/src/github.com/spacelavr/pandora/pkg/client/client.go",
	}}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, c.expected, filepath.PKG(c.path))
		})
	}
}
