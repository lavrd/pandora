package rsa_test

import (
	"crypto/rsa"
	"testing"

	ursa "github.com/spacelavr/pandora/pkg/utils/rsa"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*rsa.PrivateKey, *rsa.PublicKey) {

	private, public, err := ursa.GenerateKeys()
	assert.NoError(t, err)
	assert.NotNil(t, private)
	assert.NotNil(t, public)

	return private, public
}

func TestGenerateKeys(t *testing.T) {

	t.Parallel()

	_, _ = setup(t)
}

func TestSignPSS(t *testing.T) {

	t.Parallel()

	private, _ := setup(t)

	signature, err := ursa.SignPSS(private)
	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestVerifyPSS(t *testing.T) {

	t.Parallel()

	private, public := setup(t)

	signature, err := ursa.SignPSS(private)
	assert.NoError(t, err)
	assert.NotNil(t, signature)

	cases := []struct {
		error     error
		signature []byte
	}{{
		error:     nil,
		signature: signature,
	}, {
		error:     rsa.ErrVerification,
		signature: []byte("signature"),
	}}

	for _, c := range cases {
		err = ursa.VerifyPSS(public, c.signature)
		assert.Equal(t, c.error, err)
	}
}
