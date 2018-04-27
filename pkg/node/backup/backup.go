package backup

import (
	"github.com/spacelavr/pandora/pkg/types"
)

// Backup
type Backup struct {
	blockchain types.Blockchain
}

// Backup backup blockchain
func (b *Backup) Backup() error {
	return nil
}

// Recover recover blockchain
func (b *Backup) Recover() (types.Blockchain, error) {
	return nil, nil
}
