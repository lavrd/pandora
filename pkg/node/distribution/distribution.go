package distribution

import (
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/node/rpc"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
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
	// var (
	// 	stg = env.GetStorage()
	// 	rt  = env.GetRuntime()
	// )
	//
	// recipient, err := stg.AccountFetchByPublic(*opts.PublicKey)
	// if err != nil {
	// 	return err
	// }
	// if recipient == nil {
	// 	return errors.DocumentNotFound
	// }
	//
	// cert := &types.Certificate{
	// 	Id: generator.UUID(),
	// 	Meta: &types.CertificateMeta{
	// 		Title:       *opts.Title,
	// 		Description: *opts.Description,
	// 		Date:        time.Now().UTC(),
	// 	},
	// 	Issuer: &types.Issuer{
	// 		Meta: &types.IssuerMeta{
	// 			// Name: rt.FullName,
	// 		},
	// 	},
	// 	Recipient: &types.Recipient{
	// 		Meta: &types.RecipientMeta{
	// 			Name: recipient.Meta.FullName,
	// 		},
	// 	},
	// }
	//
	// hash := sha256.SumBytes(cert.Bytes())
	//
	// ISign := ed25519.Sign(hash, rt.PrivateKey)
	// RSign := ed25519.Sign(hash, recipient.Secure.PrivateKey)
	//
	// cert.Issuer.Signature = ISign
	// cert.Recipient.Signature = RSign
	//
	// block := rt.PrepareBlock(cert, rt.Last())
	//
	// if err := stg.BlockSave(block); err != nil {
	// 	return err
	// }
	//
	// if err := stg.CertificateSave(cert); err != nil {
	// 	return err
	// }
	//
	// events.PublishNBlock(block)

	return nil
}
