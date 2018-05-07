package nacl

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/spacelavr/pandora/pkg/config"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/nacl/sign"
)

const (
	SecretKeySize  = 32
	NonceSize      = 24
	PublicKeySize  = 32
	PrivateKeySize = 64
)

func key() [SecretKeySize]byte {
	var (
		key [SecretKeySize]byte
	)

	keyBytes, _ := hex.DecodeString(config.Viper.Secure.Key)

	copy(key[:], keyBytes)

	return key
}

func Seal(data []byte) []byte {
	var (
		key   = key()
		nonce [NonceSize]byte
	)

	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], data, &nonce, &key)

	return encrypted
}

func Open(encrypted []byte) []byte {
	var (
		key   = key()
		nonce [NonceSize]byte
	)

	copy(nonce[:], encrypted[:NonceSize])

	decrypted, ok := secretbox.Open(nil, encrypted[NonceSize:], &nonce, &key)
	if !ok {
		return nil
	}

	return decrypted
}

func GenerateKeys() (*[PublicKeySize]byte, *[PrivateKeySize]byte) {
	public, private, _ := sign.GenerateKey(rand.Reader)
	return public, private
}

func Sign(msg string, key *[PrivateKeySize]byte) string {
	signature := sign.Sign(nil, []byte(msg), key)
	return fmt.Sprintf("%x", signature)
}

func Verify(signature string, key *[PublicKeySize]byte) bool {
	buf, _ := hex.DecodeString(signature)

	_, ok := sign.Open(nil, buf, key)
	return ok
}
