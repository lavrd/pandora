package crypto

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/argon2"
)

const (
	time    = 3
	memory  = 32 * 2014
	threads = 4
	keyLen  = 32

	saltSize = 256
)

func Argon2Key(key []byte) []byte {
	var (
		nonce = [saltSize]byte{}
		salt  []byte
	)

	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	copy(salt, nonce[:])

	return argon2.Key(key, salt, time, memory, threads, keyLen)
}
