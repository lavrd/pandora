package bcrypt_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/crypto/bcrypt"
	"github.com/stretchr/testify/assert"
	bc "golang.org/x/crypto/bcrypt"
)

const (
	password = "password"
)

func setup(t *testing.T) string {
	hash, err := bcrypt.Encode(password)
	assert.NoError(t, err)
	return hash
}

func TestEncode(t *testing.T) {
	_ = setup(t)
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name     string
		hash     string
		password string
		error    error
	}{
		{
			"valid password",
			setup(t),
			password,
			nil,
		},
		{
			"invalid password",
			setup(t),
			"",
			bc.ErrMismatchedHashAndPassword,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := bcrypt.Validate(c.hash, c.password)
			assert.Equal(t, c.error, err)
		})
	}
}
