package storage

import (
	"fmt"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

const (
	CAccount = "CAccount"
)

func (s *Storage) AccountFetchByEmail(email string) (*pb.Member, error) {
	var (
		acc   = &pb.Member{}
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
func (s *Storage) AccountFetchByPublic(public string) (*pb.Member, error) {
	var (
		acc   = &pb.Member{}
		query = fmt.Sprintf(
			"for a in %s filter a.public_key.public_key == @public return a",
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
func (s *Storage) AccountSave(acc *pb.Member) error {
	_, err := s.Write(CAccount, acc)
	if err != nil {
		// todo need handle error that already exists or not? check this level up
		return err
	}

	return nil
}
