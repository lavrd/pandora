package distribution

import (
	"time"

	"github.com/spacelavr/pandora/pkg/api/routes/request"
	"github.com/spacelavr/pandora/pkg/types"
)

// CertificateIssue issue certificate
func (d *Distribution) CertificateIssue(opts *request.CertificateIssue) (*types.Certificate, error) {
	cert := &types.Certificate{
		Meta: &types.CertificateMeta{
			Title: *opts.Title,
			Date:  time.Now().UTC(),
		},
		Issuer:    &types.Issuer{},
		Recipient: &types.Recipient{},
	}

	if opts.Description != nil {
		cert.Meta.Description = *opts.Description
	}

	return cert, nil
}
