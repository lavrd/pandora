package argon2

import (
	"github.com/spacelavr/pandora/pkg/config"
	"golang.org/x/crypto/argon2"
)

const (
	time    = 3
	memory  = 32 * 2014
	threads = 4
	keyLen  = 32
)

func Key(key []byte) []byte {
	return argon2.Key(key, []byte(config.Viper.Secure.Salt), time, memory, threads, keyLen)
}
