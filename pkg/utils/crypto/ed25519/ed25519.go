package ed25519

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/ed25519"
)

func Sign(data []byte, key ed25519.PrivateKey) string {
	return fmt.Sprintf("%x", ed25519.Sign(key, data))
}

func GenerateKeys() (ed25519.PublicKey, ed25519.PrivateKey) {
	public, private, _ := ed25519.GenerateKey(rand.Reader)
	return public, private
}
