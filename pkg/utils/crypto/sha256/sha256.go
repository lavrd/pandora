package sha256

import (
	"crypto/sha256"
	"fmt"
)

// Compute computing p hash sum
func Compute(p string) string {
	h := sha256.New()
	h.Write([]byte(p))
	sum := h.Sum(nil)

	return fmt.Sprintf("%x", sum)
}
