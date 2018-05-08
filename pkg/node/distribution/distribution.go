package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/ed25519"
	"github.com/spacelavr/pandora/pkg/utils/crypto/sha256"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/generator"
)

func CertificateIssue(opts *request.CertificateIssue) error {
	var (
		stg = env.GetStorage()
		rt  = env.GetRuntime()
	)

	recipient, err := stg.AccountFetchByPublic(*opts.PublicKey)
	if err != nil {
		return err
	}
	if recipient == nil {
		return errors.DocumentNotFound
	}

	cert := &types.Certificate{
		Id: generator.UUID(),
		Meta: &types.CertificateMeta{
			Title:       *opts.Title,
			Description: *opts.Description,
			Date:        time.Now().UTC(),
		},
		Issuer: &types.Issuer{
			Meta: &types.IssuerMeta{
				// Name: rt.FullName,
			},
		},
		Recipient: &types.Recipient{
			Meta: &types.RecipientMeta{
				Name: recipient.Meta.FullName,
			},
		},
	}

	hash := sha256.SumBytes(cert.Bytes())

	ISign := ed25519.Sign(hash, rt.PrivateKey)
	RSign := ed25519.Sign(hash, recipient.Secure.PrivateKey)

	cert.Issuer.Signature = ISign
	cert.Recipient.Signature = RSign

	block := rt.PrepareBlock(cert, rt.Last())

	if err := stg.BlockSave(block); err != nil {
		return err
	}

	if err := stg.CertificateSave(cert); err != nil {
		return err
	}

	events.PublishNBlock(block)

	return nil
}
