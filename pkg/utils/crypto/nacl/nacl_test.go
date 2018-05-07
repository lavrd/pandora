package nacl_test

import (
	"encoding/pem"
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/crypto/nacl"
)

func TestTest(t *testing.T) {
	public, private := nacl.GenerateKeys()
}
