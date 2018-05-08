package types

import (
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

type (
	ResAccountFetch struct {
		Acc *Account
		Err *errors.Response
	}
)
