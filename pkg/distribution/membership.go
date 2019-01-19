package distribution

import (
	"encoding/hex"

	"pandora/pkg/membership/env"
	"pandora/pkg/pb"
	"pandora/pkg/storage/arangodb"
	"pandora/pkg/utils/crypto/ed25519"
	"pandora/pkg/utils/crypto/sha256"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/mail"
)

// Membership
type Membership struct {
	storage *arangodb.Arangodb
}

// ConfirmMember confirm member
func (m *Membership) ConfirmMember(candidate *pb.MemberMeta) (*pb.PublicKey, error) {
	mem, err := m.storage.MemberFetchByEmail(candidate.Email)
	if err != nil && err != errors.ErrNotFound {
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

	if err := m.storage.MemberSave(mem); err != nil {
		return nil, err
	}

	if err = mail.SendCredentials(candidate.Email, mem.PublicKey); err != nil {
		return nil, err
	}

	return mem.PublicKey, nil
}

// MemberFetch fetch member
func (m *Membership) MemberFetch(key *pb.PublicKey) (*pb.Member, error) {
	return m.storage.MemberFetchByPublic(key)
}

// SignCert sign cert
func (m *Membership) SignCert(cert *pb.Cert) (*pb.Cert, error) {
	recipient, err := m.storage.MemberFetchByPublic(cert.Recipient.PublicKey)
	if err != nil {
		return nil, err
	}

	issuer, err := m.storage.MemberFetchByPublic(cert.Issuer.PublicKey)
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

// New returns new membership distribution
func NewMembership() *Membership {
	return &Membership{storage: env.GetStorage()}
}
