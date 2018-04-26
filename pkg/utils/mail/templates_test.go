package mail_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/utils/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendAccountCreated(t *testing.T) {
	config.Viper.Mail.Send = false
	config.Viper.Mail.Templates.Account.Created = "./templates/account/created.html"

	err := mail.SendAccountCreated("", "")
	assert.NoError(t, err)
}

func TestSendAccountRecovery(t *testing.T) {
	config.Viper.Mail.Send = false
	config.Viper.Mail.Templates.Account.Recovery = "./templates/account/recovery.html"

	err := mail.SendAccountRecovery("", "")
	assert.NoError(t, err)
}
