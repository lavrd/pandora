package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/generator"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

type Distribution struct{}

func (d *Distribution) Candidate(opts *request.Candidate) error {
	if err := rpc.Register(&pb.Candidate{
		Name:  *opts.Name,
		Email: *opts.Email,
	}); err != nil {
		return err
	}
	return nil
}

func (d *Distribution) FetchAccount(opts *request.AccountFetch) (*pb.Member, error) {
	acc, err := rpc.FetchAccount(&pb.PublicKey{
		PublicKey: *opts.PublicKey,
	})
	if err != nil {
		return nil, err
	}

	log.Debug(acc)

	// todo added convert pkg
	return &pb.Member{
		PublicKey: acc.PublicKey,
		Meta: &pb.MemberMeta{
			Email: acc.Meta.Email,
			Name:  acc.Meta.Name,
		},
	}, nil
}

func CertificateIssue(opts *request.CertificateIssue) error {
	var (
		rt = env.GetRuntime()
	)

	if err := rpc.Issue(&pb.Cert{
		Id: generator.Id(),
		Meta: &pb.CertMeta{
			Timestamp:   time.Now().UTC().UnixNano() / 1000000,
			Description: *opts.Description,
			Title:       *opts.Title,
		},
		Recipient: &pb.Participant{
			PublicKey: &pb.PublicKey{
				PublicKey: *opts.PublicKey,
			},
		},
		Issuer: &pb.Participant{
			PublicKey: &pb.PublicKey{
				PublicKey: rt.PublicKey,
			},
		},
	}); err != nil {
		return err
	}

	return nil
}
