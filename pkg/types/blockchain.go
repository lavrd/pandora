package types

import (
	"fmt"
	"time"
)

// Block
type Block struct {
	Index     int          `json:"index"`
	Hash      string       `json:"hash"`
	PrevHash  string       `json:"prev_hash"`
	Timestamp time.Time    `json:"timestamp"`
	Cert      *Certificate `json:"cert,omitempty"`
}

// String returns block as a string
// if only fmt.Sprint -> stack overflow
func (b Block) String() string {
	return fmt.Sprintf("%s%s%d%s%v",
		b.Hash,
		b.PrevHash,
		b.Index,
		b.Timestamp.String(),
		b.Cert,
	)
}

// Blockchain
type Blockchain []*Block
