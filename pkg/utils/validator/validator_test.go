package validator_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/validator"
	"github.com/stretchr/testify/assert"
)

func TestIsEmail(t *testing.T) {
	cases := []struct {
		email string
		valid bool
	}{
		{
			"pandora@doesnotexists.system",
			true,
		},
		{
			"pandora@doesnotexists.",
			false,
		},
		{
			"pandora@doesnotexistssystem",
			false,
		},
		{
			"@doesnotexists.system",
			false,
		},
		{
			"pandoradoesnotexists.system",
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.email, func(t *testing.T) {
			assert.Equal(t, c.valid, validator.IsEmail(c.email))
		})
	}
}

func TestIsPassword(t *testing.T) {
	cases := []struct {
		password string
		valid    bool
	}{
		{
			"1111111111",
			true,
		},
		{
			"111",
			false,
		},
		{
			"11111111111111111111111111111111111111111111111111111111111111111",
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.password, func(t *testing.T) {
			assert.Equal(t, c.valid, validator.IsPassword(c.password))
		})
	}
}
