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

func (b *Block) Bytes() []byte {
	return []byte(fmt.Sprint(b))
}

// Blockchain
type Blockchain []*Block
