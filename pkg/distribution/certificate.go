package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/api/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

// CertIssue issue certificate
func (d *Distribution) CertificateIssue(opts *request.CertificateIssue) (*types.Certificate, error) {
	issuer, err := d.AccountFetch(*opts.IssuerEmail)
	if err != nil {
		return nil, err
	}
	if issuer == nil {
		return nil, errors.AccountNotFound
	}

	recipient, err := d.AccountFetch(*opts.RecipientEmail)
	if err != nil {
		return nil, err
	}
	if recipient == nil {
		return nil, errors.AccountNotFound
	}

	cert := &types.Certificate{
		Meta: &types.CertificateMeta{
			Title:       *opts.Title,
			Description: *opts.Description,
			Date:        time.Now().UTC(),
		},
		Issuer: &types.Issuer{
			PublicKey: issuer.PublicKey,
			Meta: &types.IssuerMeta{
				Name: issuer.Meta.Name,
			},
		},
		Recipient: &types.Recipient{
			PublicKey: recipient.PublicKey,
			Meta: &types.RecipientMeta{
				Name: recipient.Meta.Name,
			},
		},
	}

	if err = d.CertSave(cert); err != nil {
		return nil, err
	}

	return cert, nil
}
