package mail

import (
	"bytes"
	"html/template"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func execute(email, subject, path string, data interface{}) error {
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

	return send(email, subject, buf.String())
}

func SendCredentials(email, publicKey string) error {
	data := &struct {
		PublicKey string
	}{
		PublicKey: publicKey,
	}

	return execute(
		email,
		config.Viper.Mail.Subjects.Credentials,
		config.Viper.Mail.Templates.Credentials,
		data,
	)
}
