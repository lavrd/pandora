package distribution

import (
	"github.com/spacelavr/pandora/pkg/membership/env"
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

func New() *Distribution {
	return &Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}
}

func (d *Distribution) AcceptCandidate(candidate *pb.Candidate) (*pb.PublicKey, error) {
	acc, err := d.Storage.AccountFetchByEmail(candidate.Email)
	if err != nil {
		return nil, err
	}
	if acc != nil {
		return &pb.PublicKey{PublicKey: acc.PublicKey}, errors.AlreadyExists
	}

	acc = d.Runtime.AcceptCandidate(candidate)

	if err := d.Storage.AccountSave(acc); err != nil {
		return nil, err
	}

	if err = mail.SendCredentials(candidate.Email, acc.PublicKey); err != nil {
		return nil, err
	}

	return &pb.PublicKey{PublicKey: acc.PublicKey}, nil
}

func (d *Distribution) AccountFetch(key *pb.PublicKey) (*pb.Account, error) {
	acc, err := d.Storage.AccountFetchByPublic(key.PublicKey)
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

func (d *Distribution) Issue(cert *pb.Cert) (*pb.Cert, error) {
	recipient, err := d.Storage.AccountFetchByPublic(cert.Recipient.PublicKey.PublicKey)
	if err != nil {
		return nil, err
	}
	if recipient == nil {
		return nil, errors.NotFound
	}

	issuer, err := d.Storage.AccountFetchByPublic(cert.Issuer.PublicKey.PublicKey)
	if err != nil {
		return nil, err
	}
	if issuer == nil {
		return nil, errors.NotFound
	}

	cert = d.Runtime.SignCert(cert, issuer.Secure.PrivateKey, recipient.Secure.PrivateKey)

	cert.Recipient.Meta = &pb.RecipientMeta{
		Name: recipient.Meta.FullName,
	}

	cert.Issuer.Meta = &pb.IssuerMeta{
		Name: issuer.Meta.FullName,
	}

	return cert, nil
}
