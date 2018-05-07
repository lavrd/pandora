package nacl_test

import (
	"fmt"
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/crypto/nacl"
)

func TestTest(t *testing.T) {
	public, private := nacl.GenerateKeys()

	fmt.Println(public)
	fmt.Println(private)

	signature := nacl.Sign("hello", private)
	fmt.Println(signature)

	fmt.Println(nacl.Verify(signature, public))
}
