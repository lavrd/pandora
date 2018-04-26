package generator

import (
	"github.com/satori/go.uuid"
	"github.com/spacelavr/pandora/pkg/types"
)

// UUID generate uuid string
func UUID() string {
	return uuid.NewV4().String()
}

// Password generate password from uuid
func Password() string {
	return UUID()[:types.MinPasswordLen]
}
