package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/utils/errors"
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

// Marshal marshal rsa keys
func Marshal(pri *rsa.PrivateKey, pub *rsa.PublicKey) (string, string) {
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pri),
	}
	private := pem.EncodeToMemory(block)

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pub),
	}
	public := pem.EncodeToMemory(block)

	return string(private), string(public)
}

// Unmarshal unmarshal rsa keys
func Unmarshal(pri, pub string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pri))
	if block == nil {
		log.Error(errors.PemBlockParseFailed)
		return nil, nil, errors.PemBlockParseFailed
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	block, _ = pem.Decode([]byte(pub))
	if block == nil {
		log.Error(errors.PemBlockParseFailed)
		return nil, nil, errors.PemBlockParseFailed
	}
	public, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	return private, public, nil
}

// SignPSS calculates signature
func SignPSS(key *rsa.PrivateKey) (string, error) {
	signature, err := rsa.SignPSS(rand.Reader, key, hash, hashed, opts)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return fmt.Sprintf("%x", signature), nil
}

// VerifyPSS verify signature
func VerifyPSS(key *rsa.PublicKey, signature string) error {
	err := rsa.VerifyPSS(key, hash, hashed, []byte(signature), opts)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
