package ed25519

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/ed25519"
)

const (
	PublicKeySize = ed25519.PublicKeySize
)

type (
	PrivateKey = ed25519.PrivateKey
	PublicKey = ed25519.PublicKey
)

func Sign(data []byte, key PrivateKey) string {
	return fmt.Sprintf("%x", ed25519.Sign(key, data))
}

func Verify(data []byte, signature string, key PublicKey) bool {
	buf, _ := hex.DecodeString(signature)
	return ed25519.Verify(key, data, buf)
}

func GenerateKeys() (PublicKey, PrivateKey) {
	public, private, _ := ed25519.GenerateKey(rand.Reader)
	return public, private
}
