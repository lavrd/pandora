package generator

import (
	"github.com/satori/go.uuid"
)

// todo rename
func UUID() string {
	v4, _ := uuid.NewV4()
	return v4.String()
}
