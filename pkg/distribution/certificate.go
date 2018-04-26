package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/api/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/rsa"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/generator"
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
	if recipient.Meta.Type != types.TypeRecipient {
		return nil, errors.IssueCertToNonRecipient
	}

	IPrivate, _, err := rsa.Unmarshal(issuer.Secure.PrivateKey, issuer.PublicKey)
	if err != nil {
		return nil, err
	}

	RPrivate, _, err := rsa.Unmarshal(recipient.Secure.PrivateKey, recipient.PublicKey)
	if err != nil {
		return nil, err
	}

	ISign, err := rsa.SignPSS(IPrivate)
	if err != nil {
		return nil, err
	}

	RSign, err := rsa.SignPSS(RPrivate)
	if err != nil {
		return nil, err
	}

	cert := &types.Certificate{
		Id: generator.UUID(),
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
			Signature: ISign,
		},
		Recipient: &types.Recipient{
			PublicKey: recipient.PublicKey,
			Meta: &types.RecipientMeta{
				Name: recipient.Meta.Name,
			},
			Signature: RSign,
		},
	}

	if err = d.CertificateSave(cert); err != nil {
		return nil, err
	}

	return cert, nil
}

// CertificateView returns certificate
func (d *Distribution) CertificateView(id string) (*types.Certificate, error) {
	cert, err := d.CertificateFetch(id)
	if err != nil {
		return nil, err
	}
	if cert == nil {
		return nil, errors.CertificateNotFound
	}

	return cert, nil
}
