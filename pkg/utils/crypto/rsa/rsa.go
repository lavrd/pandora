package rsa

// why without (_ "crypto/sha256") import sha256 is unavailable?
import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"

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
	hash := hash
	pssh := hash.New()
	pssh.Write([]byte(viper.GetString("secure.PSSMessage")))
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

// SignPSS calculates signature
func SignPSS(key *rsa.PrivateKey) ([]byte, error) {
	signature, err := rsa.SignPSS(rand.Reader, key, hash, hashed, opts)
	if err != nil {
		log.Error(err)
	}
	return signature, err
}

// VerifyPSS verify signature
func VerifyPSS(key *rsa.PublicKey, signature []byte) error {
	err := rsa.VerifyPSS(key, hash, hashed, signature, opts)
	if err != nil {
		log.Error(err)
	}
	return err
}
