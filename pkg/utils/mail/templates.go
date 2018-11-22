package mail

import (
	"bytes"
	"html/template"

	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
)

func execute(to, subject, path string, data interface{}) error {
	var buf bytes.Buffer

	tpl, err := template.ParseFiles(path)
	if err != nil {
		return errors.WithStack(err)
	}

	if err = tpl.Execute(&buf, data); err != nil {
		return errors.WithStack(err)
	}

	return send(to, subject, buf.String())
}

// SendCertificate send certificate mail
func SendCertificate(to string, cert *pb.Cert) error {
	const (
		Subject  = "Certificate"
		Template = "./pkg/utils/mail/templates/certificate.html"
	)

	return execute(
		to,
		Subject,
		Template,
		cert,
	)
}

// SendCredentials send credentials mail
func SendCredentials(to string, publicKey *pb.PublicKey) error {
	const (
		Subject  = "Credentials"
		Template = "./pkg/utils/mail/templates/credentials.html"
	)

	return execute(
		to,
		Subject,
		Template,
		publicKey,
	)
}
