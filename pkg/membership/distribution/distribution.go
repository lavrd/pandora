package distribution

import (
	"github.com/spacelavr/pandora/pkg/membership/runtime"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/mail"
)

type Distribution struct {
	*runtime.Runtime
	*storage.Storage
}

func (d *Distribution) AcceptCandidate(candidate *pb.Candidate) (*pb.PublicKey, error) {
	acc, err := d.AccountFetchByEmail(candidate.Email)
	if err != nil {
		return nil, err
	}
	if acc != nil {
		return nil, errors.AlreadyExists
	}

	acc = d.Runtime.AcceptCandidate(candidate)

	if err := d.AccountSave(acc); err != nil {
		return nil, err
	}

	if err = mail.SendCredentials(candidate.Email, string(acc.PublicKey)); err != nil {
		return nil, err
	}

	return &pb.PublicKey{PublicKey: acc.PublicKey}, nil
}

func (d *Distribution) AccountFetch(key *pb.PublicKey) (*pb.Account, error) {
	acc, err := d.AccountFetchByPublic(key.PublicKey)
	if err != nil {
		return nil, err
	}
	if acc == nil {
		return nil, errors.NotFound
	}

	return &pb.Account{
		PublicKey: acc.PublicKey,
		Meta: &pb.AccountMeta{
			FullName: acc.Meta.FullName,
			Email:    acc.Meta.Email,
		},
	}, nil
}
