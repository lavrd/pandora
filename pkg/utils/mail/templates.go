package mail

import (
	"bytes"
	"html/template"

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

// SendAccountRecovery send account recovery email
func SendAccountRecovery(email, password string) error {
	// data := &types.AccountRecovery{Password: password}
	// return execute(
	// 	email,
	// 	config.Viper.Mail.Subjects.Account.Recovery,
	// 	config.Viper.Mail.Templates.Account.Recovery,
	// 	data,
	// )
	return nil
}

// SendAccountCreated send account created email
func SendAccountCreated(email, public string) error {
	// data := &types.Account{Secure: &types.AccountSecure{Password: password}}
	// return execute(
	// 	email,
	// 	config.Viper.Mail.Subjects.Account.Created,
	// 	config.Viper.Mail.Templates.Account.Created,
	// 	data,
	// )
	return nil
}
