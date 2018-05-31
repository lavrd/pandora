package mail

import (
	"bytes"
	"html/template"

	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func execute(to, subject, path string, data interface{}) error {
	var buf bytes.Buffer

	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Error(err)
		return err
	}

	if err = tpl.Execute(&buf, data); err != nil {
		log.Error(err)
		return err
	}

	return send(to, subject, buf.String())
}

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
