package types

import (
	"golang.org/x/crypto/ed25519"
)

type (
	Session struct {
		Token string `json:"token"`
	}

	Account struct {
		PublicKey ed25519.PublicKey `json:"public_key"`
		Secure    *AccountSecure    `json:"secure,omitempty"`
		Meta      *AccountMeta      `json:"meta"`
	}

	AccountMeta struct {
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}

	AccountSecure struct {
		PrivateKey ed25519.PrivateKey `json:"private_key"`
	}
)

// Public returns public account info
func (acc *Account) Public() *Account {
	acc.Secure = nil
	return acc
}
