package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
)

const (
	BucketAccount = "account"
)

// AccountFetch fetch account from storage
func (s *Storage) AccountFetch(email string) (*types.Account, error) {
	acc := &types.Account{}

	err := s.Get(BucketAccount, email, acc)
	if err != nil {
		return nil, err
	}

	if (types.Account{}) == *acc {
		return nil, nil
	}

	return acc, nil
}

// AccountPut put account to storage
func (s *Storage) AccountPut(acc *types.Account) error {
	return s.Put(BucketAccount, acc.Email, acc)
}
