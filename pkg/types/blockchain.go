package types

import (
	"fmt"
	"time"
)

type (
	CertBlock struct {
		*Block
	}

	MasterBlock struct {
		*Block
		PublicKey []byte    `json:"public_key"`
		CertChain CertChain `json:"cert_chain"`
	}

	Block struct {
		Key       string    `json:"_key"`
		Index     int       `json:"index"`
		Hash      string    `json:"hash"`
		PrevHash  string    `json:"prev_hash"`
		Timestamp time.Time `json:"timestamp"`
	}

	CertChain []*CertBlock
	MasterChain []*MasterBlock
)

func (b *MasterBlock) Bytes() []byte {
	return []byte(fmt.Sprint(b))
}

func (b *CertBlock) Bytes() []byte {
	return []byte(fmt.Sprint(b))
}
