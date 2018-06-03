package mail_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendCredentials(t *testing.T) {
	conf.Viper.Membership.Mail.Send = false

	err := mail.SendCredentials("", &pb.PublicKey{})
	assert.NoError(t, err)
}

func TestSendCertificate(t *testing.T) {
	conf.Viper.Membership.Mail.Send = false

	err := mail.SendCertificate("", &pb.Cert{})
	assert.NoError(t, err)
}
