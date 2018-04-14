package generator

import (
	"github.com/satori/go.uuid"
	"github.com/spacelavr/pandora/pkg/types"
)

// Password generate password from uuid
func Password() string {
	return uuid.NewV4().String()[:types.MinPasswordLen]
}
