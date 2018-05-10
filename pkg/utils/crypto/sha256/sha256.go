package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func sum(p []byte) []byte {
	h := sha256.New()
	h.Write(p)
	sum := h.Sum(nil)
	return sum
}

func SumString(p []byte) string {
	return hex.EncodeToString(sum(p))
}

func SumBytes(p []byte) []byte {
	return sum(p)
}
