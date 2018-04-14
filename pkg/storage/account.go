package storage

import (
	"github.com/spacelavr/pandora/pkg/types"
)

// AccountFetch fetch account from storage
func AccountFetch(email string) (*types.Account, error) {
	acc := &types.Account{}
	err := get(BucketAccount, email, acc)

	if (types.Account{}) == *acc {
		return nil, nil
	}
	return acc, err
}

// AccountPut put account to storage
func AccountPut(acc *types.Account) error {
	return put(BucketAccount, acc.Email, acc)
}
