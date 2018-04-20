package jwt_test

import (
	"testing"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/jwt"
	"github.com/stretchr/testify/assert"
)

var (
	email = "asd@asd.asd"
)

func setup(t *testing.T) string {
	token, err := jwt.New(&types.Account{Meta: &types.AccountMeta{Email: email}})
	assert.NoError(t, err)

	return token
}

func TestGenerate(t *testing.T) {
	_ = setup(t)
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name  string
		token string
		email string
		error error
	}{
		{
			"valid token",
			setup(t),
			email,
			nil,
		},
		{
			"invalid token",
			"token",
			"",
			jwtGo.NewValidationError(
				"token contains an invalid number of segments",
				jwtGo.ValidationErrorMalformed,
			),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			e, err := jwt.Validate(c.token)
			assert.Equal(t, c.error, err)
			assert.Equal(t, c.email, e)
		})
	}
}
