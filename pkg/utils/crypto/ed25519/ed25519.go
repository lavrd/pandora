package ed25519

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ed25519"
)

func Sign(data []byte, key ed25519.PrivateKey) string {
	return fmt.Sprintf("%x", ed25519.Sign(key, data))
}

func Verify(data []byte, signature string, key ed25519.PublicKey) bool {
	buf, _ := hex.DecodeString(signature)
	return ed25519.Verify(key, data, buf)
}

func GenerateKeys() (ed25519.PublicKey, ed25519.PrivateKey) {
	public, private, _ := ed25519.GenerateKey(rand.Reader)
	return public, private
}
