package runtime

import (
	"encoding/hex"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/ed25519"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
)

type Runtime struct{}

func New() *Runtime {
	return &Runtime{}
}

func (r *Runtime) AcceptCandidate(candidate *pb.Candidate) *types.Account {
	publicKey, privateKey := ed25519.GenerateKeys()

	return &types.Account{
		Key: hex.EncodeToString(publicKey),
		Meta: &types.AccountMeta{
			Email:    candidate.Email,
			FullName: candidate.FullName,
		},
		PublicKey: hex.EncodeToString(publicKey),
		Secure: &types.AccountSecure{
			PrivateKey: privateKey,
		},
	}
}

func (r *Runtime) SignCert(cert *pb.Cert, issuerKey []byte, recipientKey []byte) *pb.Cert {
	buf, _ := hex.DecodeString(cert.String())
	hash := sha256.SumBytes(buf)

	ISignature := ed25519.Sign(hash, issuerKey)
	RSignature := ed25519.Sign(hash, recipientKey)

	cert.Issuer.Signature = ISignature
	cert.Recipient.Signature = RSignature

	return cert
}
