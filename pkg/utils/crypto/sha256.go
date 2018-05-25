package crypto

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

func SumString(p string) string {
	buf, _ := hex.DecodeString(p)
	return hex.EncodeToString(sum(buf))
}
