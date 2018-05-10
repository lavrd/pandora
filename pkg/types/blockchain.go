package types

import (
	"encoding/hex"
	"fmt"
	"time"
)

type (
	CertBlock struct {
		*Block
		PublicKey string `json:"public_key"`
	}

	MasterBlock struct {
		*Block
		PublicKey string    `json:"public_key"`
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
	buf, _ := hex.DecodeString(fmt.Sprint(b))
	return buf
}

func (b *CertBlock) Bytes() []byte {
	buf, _ := hex.DecodeString(fmt.Sprint(b))
	return buf
}
