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

func (d *Distribution) CandidateCheck(candidate *pb.Candidate) error {
	acc, err := d.AccountFetchByEmail(candidate.Email)
	if err != nil {
		return err
	}
	if acc != nil {
		return errors.AlreadyExists
	}

	acc = d.AcceptCandidate(candidate)

	if err := d.AccountSave(acc); err != nil {
		return err
	}

	if err = mail.SendCredentials(candidate.Email, string(acc.PublicKey)); err != nil {
		return err
	}

	return nil
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
