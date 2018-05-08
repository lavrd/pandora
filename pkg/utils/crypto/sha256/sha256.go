package sha256

import (
	"crypto/sha256"
	"fmt"
)

func sum(p []byte) []byte {
	h := sha256.New()
	h.Write(p)
	sum := h.Sum(nil)
	return sum
}

func SumString(p []byte) string {
	return fmt.Sprintf("%x", sum(p))
}

func SumBytes(p []byte) []byte {
	return sum(p)
}
