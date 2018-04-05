package filepath_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/filepath"
	"github.com/stretchr/testify/assert"
)

func TestPKG(t *testing.T) {

	t.Parallel()

	cases := []struct {
		expected string
		path     string
	}{{
		expected: "/pkg/utils/filepath/filepath.go",
		path:     "/opt/go/src/github.com/spacelavr/pandora/pkg/utils/filepath/filepath.go",
	}, {
		expected: "/pkg/utils/http/http.go",
		path:     "/opt/go/src/github.com/spacelavr/pandora/pkg/utils/http/http.go",
	}, {
		expected: "/pkg/client/client.go",
		path:     "/opt/go/src/github.com/spacelavr/pandora/pkg/client/client.go",
	}}

	for _, c := range cases {
		assert.Equal(t, c.expected, filepath.PKG(c.path))
	}
}
