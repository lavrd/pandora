package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/generator"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

type Distribution struct{}

func (d *Distribution) Candidate(opts *request.Candidate) error {
	if err := rpc.Register(&pb.Candidate{
		FullName: *opts.FullName,
		Email:    *opts.Email,
	}); err != nil {
		return err
	}
	return nil
}

func (d *Distribution) FetchAccount(opts *request.AccountFetch) (*types.Account, error) {
	acc, err := rpc.FetchAccount(&pb.PublicKey{
		PublicKey: *opts.PublicKey,
	})
	if err != nil {
		return nil, err
	}

	log.Debug(acc)

	// todo added convert pkg
	return &types.Account{
		PublicKey: acc.PublicKey,
		Meta: &types.AccountMeta{
			Email:    acc.Meta.Email,
			FullName: acc.Meta.FullName,
		},
	}, nil
}

func CertificateIssue(opts *request.CertificateIssue) error {
	var (
		stg = env.GetStorage()
		rt  = env.GetRuntime()
	)

	return nil
}
