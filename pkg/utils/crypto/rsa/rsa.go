package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

var (
	opts = &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
	}
	hash            = crypto.SHA256
	hashed          []byte
	ErrVerification = errors.New("ErrVerification")
)

func init() {
	hash := hash
	pssh := hash.New()
	pssh.Write([]byte(config.Viper.Secure.PSSMessage))
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

// Encode encode rsa keys
func Encode(pri *rsa.PrivateKey, pub *rsa.PublicKey) (string, string) {
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

// Decode decode key to pem block
func Decode(key string) (*pem.Block, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		log.Error(errors.PemBlockParseFailed)
		return nil, errors.PemBlockParseFailed
	}

	return block, nil
}

// DecodePrivate decode private key
func DecodePrivate(pri string) (*rsa.PrivateKey, error) {
	block, err := Decode(pri)
	if err != nil {
		return nil, err
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return private, nil
}

// DecodePublic decode public key
func DecodePublic(pub string) (*rsa.PublicKey, error) {
	block, err := Decode(pub)
	if err != nil {
		return nil, err
	}

	public, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return public, nil
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
	h, err := hex.DecodeString(signature)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := rsa.VerifyPSS(key, hash, hashed, h, opts); err != nil {
		if err == rsa.ErrVerification {
			return ErrVerification
		}
		log.Error(err)
		return err
	}

	return nil
}
