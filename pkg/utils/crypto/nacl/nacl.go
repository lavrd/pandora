package nacl

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/utils/crypto/argon2"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	SecretKeySize = 32
	NonceSize     = 24
)

func key() [SecretKeySize]byte {
	var (
		key [SecretKeySize]byte
	)

	keyBytes, _ := hex.DecodeString(config.Viper.Secure.Key)
	keyBytes = argon2.Key(keyBytes)

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
