package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
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
		rt = env.GetRuntime()
	)

	log.Debug(1)
	if err := rpc.Issue(&pb.Cert{
		Id: generator.UUID(),
		Meta: &pb.CertMeta{
			Timestamp:   time.Now().UTC().Unix(),
			Description: *opts.Description,
			Title:       *opts.Title,
		},
		Recipient: &pb.CertRecipient{
			PublicKey: &pb.PublicKey{
				PublicKey: *opts.PublicKey,
			},
		},
		Issuer: &pb.CertIssuer{
			PublicKey: &pb.PublicKey{
				PublicKey: rt.PublicKey,
			},
		},
	}); err != nil {
		return err
	}
	log.Debug(2)

	return nil
}
