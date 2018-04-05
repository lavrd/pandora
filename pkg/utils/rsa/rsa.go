package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spf13/viper"
)

var (
	opts = &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
	}
	hash   = crypto.SHA256
	hashed []byte
)

func init() {
	PSSMessage := []byte(viper.GetString("rsa.PSSMessage"))
	hash := crypto.SHA256
	pssh := hash.New()
	pssh.Write(PSSMessage)
	hashed = pssh.Sum(nil)
}

// GenerateKeys generate rsa keys
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {

	private, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	public := &private.PublicKey

	return private, public, nil
}

// SignPSS create signature
func SignPSS(key *rsa.PrivateKey) ([]byte, error) {

	signature, err := rsa.SignPSS(rand.Reader, key, hash, hashed, opts)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return signature, nil
}

// VerifyPSS verify signature
func VerifyPSS(key *rsa.PublicKey, signature []byte) error {
	// dont need error log because error only if verification failed
	return rsa.VerifyPSS(key, hash, hashed, signature, opts)
}
