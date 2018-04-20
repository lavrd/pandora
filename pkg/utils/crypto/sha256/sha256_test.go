package sha256_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	cases := []struct {
		p    string
		hash string
	}{
		{
			"well done!",
			"4908b517f8a915997b04d4805ea9aa4120e06d147217c6f32a7be36fabf1020a",
		},
		{
			"awesome!",
			"4e92163f0888d2403d61a02411a9979c711852e807f13549004e7ad147866105",
		},
	}

	for _, c := range cases {
		t.Run(c.p, func(t *testing.T) {
			assert.Equal(t, c.hash, sha256.Compute(c.p))
		})
	}
}
