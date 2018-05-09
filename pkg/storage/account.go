package storage

import (
	"fmt"

	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

const (
	CAccount = "CAccount"
)

// AccountFetchByEmail fetch account from storage by email
func (s *Storage) AccountFetchByEmail(email string) (*types.Account, error) {
	var (
		acc   = &types.Account{}
		query = fmt.Sprintf(
			"for a in %s filter a.meta.email == @email return a",
			CAccount,
		)
		vars = map[string]interface{}{
			"email": email,
		}
	)

	_, err := s.Exec(query, vars, acc)
	if err != nil {
		if err == errors.NotFound {
			return nil, nil
		}
		return nil, err
	}

	return acc, nil
}

// AccountFetchByPublic fetch account from storage by public key
func (s *Storage) AccountFetchByPublic(public string) (*types.Account, error) {
	var (
		acc   = &types.Account{}
		query = fmt.Sprintf(
			"for a in %s filter a.public_key == @public return a",
			CAccount,
		)
		vars = map[string]interface{}{
			"public": public,
		}
	)

	_, err := s.Exec(query, vars, acc)
	if err != nil {
		if err == errors.NotFound {
			return nil, nil
		}
		return nil, err
	}

	return acc, nil
}

// AccountSave save account to storage
func (s *Storage) AccountSave(acc *types.Account) error {
	_, err := s.Write(CAccount, acc)
	if err != nil {
		return err
	}

	return nil
}

// AccountUpdate update account in storage
func (s *Storage) AccountUpdate(acc *types.Account) error {
	var (
		query = fmt.Sprintf(
			"for a in %s filter a.meta.email == @email replace a with @acc in %%s",
			CAccount,
		)
		vars = map[string]interface{}{
			"email": acc.Meta.Email,
			"acc":   acc,
		}
	)

	_, err := s.Exec(query, vars, acc)
	if err != nil {
		if err == errors.NotFound {
			return nil
		}
		return err
	}

	return nil
}
