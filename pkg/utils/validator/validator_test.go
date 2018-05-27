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
			"@d.system",
			false,
		},
		{
			"p.system",
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.email, func(t *testing.T) {
			assert.Equal(t, c.valid, validator.IsEmail(c.email))
		})
	}
}
