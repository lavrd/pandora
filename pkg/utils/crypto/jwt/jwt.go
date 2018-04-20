package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spf13/viper"
)

var (
	key = []byte(viper.GetString("secure.JWTKey"))
)

// New generate new jwt token
func New(acc *types.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": acc.Meta.Email,
	})

	signed, err := token.SignedString(key)
	if err != nil {
		log.Error(err)
		return "", err
	}

	return signed, nil
}

// Validate validate jwt token and returns token email
func Validate(tkn string) (string, error) {
	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.UnexpectedSigningMethod
		}
		return key, nil
	})
	if err != nil {
		log.Error(err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"].(string), nil
	}

	return "", nil
}
