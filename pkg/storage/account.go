package storage

import (
	"github.com/arangodb/go-driver"
	"github.com/spacelavr/pandora/pkg/types"
)

const (
	CollectionAccount = "account"
)

// AccountFetch fetch account from storage
func (s *Storage) AccountFetch(email string) (*types.Account, error) {
	acc := &types.Account{}

	_, err := s.Read(CollectionAccount, email, acc)
	if err != nil {
		if driver.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return acc, nil
}

// AccountSave save account to storage
func (s *Storage) AccountSave(acc *types.Account) error {
	_, err := s.Write(CollectionAccount, acc)
	if err != nil {
		return err
	}

	return nil
}

// AccountUpdate update account in storage
func (s *Storage) AccountUpdate(acc *types.Account) error {
	_, err := s.Update(CollectionAccount, acc.Email, acc)
	if err != nil {
		return err
	}

	return nil
}
