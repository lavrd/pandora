package mail

import (
	"bytes"
	"html/template"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spf13/viper"
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

// SendAccountRecovery send account recovery email
func SendAccountRecovery(email, passowrd string) error {
	data := &types.AccountRecovery{Password: passowrd}
	return execute(
		email,
		viper.GetString("mail.subjects.account.recovery"),
		viper.GetString("mail.templates.account.recovery"),
		data,
	)
}

// SendAccountCreated send account created email
func SendAccountCreated(email, password string) error {
	data := &types.Account{Email: email, Secure: &types.AccountSecure{Password: password}}
	return execute(
		email,
		viper.GetString("mail.subjects.account.created"),
		viper.GetString("mail.templates.account.created"),
		data,
	)
}
