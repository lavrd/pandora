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

// SumString sum string by sha256 alg and returns as a string
func SumString(p string) string {
	buf, _ := hex.DecodeString(p)
	return hex.EncodeToString(sum(buf))
}

// SumString sum string by sha256 alg and returns as a bytes
func SumBytes(p string) []byte {
	buf, _ := hex.DecodeString(p)
	return sum(buf)
}
