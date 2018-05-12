package generator_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/generator"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	for i := 0; i < 10; i++ {
		assert.Equal(t, pb.MinPasswordLen, len(generator.Password()))
	}
}
