package distribution

import (
	"encoding/hex"

	"github.com/spacelavr/pandora/pkg/membership/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/storage/arangodb"
	"github.com/spacelavr/pandora/pkg/utils/crypto/ed25519"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/mail"
)

type Distribution struct {
	storage *arangodb.ArangoDB
}

func New() *Distribution {
	return &Distribution{storage: env.GetStorage()}
}

func (d *Distribution) ConfirmMember(candidate *pb.MemberMeta) (*pb.PublicKey, error) {
	mem, err := d.storage.MemberFetchByEmail(candidate.Email)
	if err != nil && err != errors.NotFound {
		return nil, err
	}
	if mem != nil {
		return mem.PublicKey, nil
	}

	publicKey, privateKey := ed25519.GenerateKeys()

	mem = &pb.Member{
		Meta: &pb.MemberMeta{
			Email: candidate.Email,
			Name:  candidate.Name,
		},
		PublicKey:  &pb.PublicKey{PublicKey: hex.EncodeToString(publicKey)},
		PrivateKey: privateKey,
	}

	if err := d.storage.MemberSave(mem); err != nil {
		return nil, err
	}

	if err = mail.SendCredentials(candidate.Email, mem.PublicKey); err != nil {
		return nil, err
	}

	return mem.PublicKey, nil
}

func (d *Distribution) MemberFetch(key *pb.PublicKey) (*pb.Member, error) {
	return d.storage.MemberFetchByPublic(key)
}

func (d *Distribution) SignCert(cert *pb.Cert) (*pb.Cert, error) {
	recipient, err := d.storage.MemberFetchByPublic(cert.Recipient.PublicKey)
	if err != nil {
		return nil, err
	}

	issuer, err := d.storage.MemberFetchByPublic(cert.Issuer.PublicKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.SumBytes(cert.String())

	ISignature := ed25519.Sign(hash, issuer.PrivateKey)
	RSignature := ed25519.Sign(hash, recipient.PrivateKey)

	cert.Issuer.Signature = ISignature
	cert.Recipient.Signature = RSignature

	cert.Recipient.Name = recipient.Meta.Name
	cert.Issuer.Name = issuer.Meta.Name

	if err := mail.SendCertificate(recipient.Meta.Email, cert); err != nil {
		return nil, err
	}

	return cert, nil
}
