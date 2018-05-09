package runtime

import (
	"github.com/spacelavr/pandora/pkg/membership/pb"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/ed25519"
)

type Runtime struct{}

func New() *Runtime {
	return &Runtime{}
}

func (r *Runtime) AcceptCandidate(candidate *pb.Candidate) *types.Account {
	publicKey, privateKey := ed25519.GenerateKeys()

	return &types.Account{
		Meta: &types.AccountMeta{
			Email:    candidate.Email,
			FullName: candidate.FullName,
		},
		PublicKey: publicKey,
		Secure: &types.AccountSecure{
			PrivateKey: privateKey,
		},
	}
}
