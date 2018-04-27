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
	issuer, err := d.AccountFetchByEmail(*opts.IssuerEmail)
	if err != nil {
		return nil, err
	}
	if issuer == nil {
		return nil, errors.DocumentNotFound
	}

	recipient, err := d.AccountFetchByEmail(*opts.RecipientEmail)
	if err != nil {
		return nil, err
	}
	if recipient == nil {
		return nil, errors.DocumentNotFound
	}
	if recipient.Meta.Type != types.TypeRecipient {
		return nil, errors.IssueCertToNonRecipient
	}

	IPrivate, err := rsa.DecodePrivate(issuer.Secure.PrivateKey)
	if err != nil {
		return nil, err
	}

	RPrivate, err := rsa.DecodePrivate(recipient.Secure.PrivateKey)
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
			Meta: &types.IssuerMeta{
				Name: issuer.Meta.Name,
			},
			Signature: ISign,
		},
		Recipient: &types.Recipient{
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
func (d *Distribution) CertificateView(opts *request.CertificateView) (*types.Certificate, error) {
	cert, err := d.CertificateFetch(*opts.Id)
	if err != nil {
		return nil, err
	}
	if cert == nil {
		return nil, errors.DocumentNotFound
	}

	return cert, nil
}

// CertificateVerify verify certificate
func (d *Distribution) CertificateVerify(opts *request.CertificateVerify) (*types.Certificate, error) {
	cert, err := d.CertificateFetch(*opts.Id)
	if err != nil {
		return nil, err
	}
	if cert == nil {
		return nil, errors.DocumentNotFound
	}

	IPublic, err := rsa.DecodePublic(*opts.IssuerPublicKey)
	if err != nil {
		return nil, err
	}

	RPublic, err := rsa.DecodePublic(*opts.RecipientPublicKey)
	if err != nil {
		return nil, err
	}

	if err := rsa.VerifyPSS(IPublic, cert.Issuer.Signature); err != nil {
		if err == rsa.ErrVerification {
			return nil, nil
		}
		return nil, err
	}

	if err := rsa.VerifyPSS(RPublic, cert.Recipient.Signature); err != nil {
		if err == rsa.ErrVerification {
			return nil, nil
		}
		return nil, err
	}

	return cert, nil
}
