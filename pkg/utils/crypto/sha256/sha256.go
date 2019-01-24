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

// CalcToString sum string by sha256 alg and returns as a string
func CalcToString(p string) string {
	buf, _ := hex.DecodeString(p)
	return hex.EncodeToString(sum(buf))
}

// CalcToBytes sum string by sha256 alg and returns as a bytes
func CalcToBytes(p string) []byte {
	buff, _ := hex.DecodeString(p)
	return sum(buff)
}
