package mail

import (
	"fmt"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"pandora/pkg/conf"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
)

func send(to, subject, html string) error {
	if !conf.Conf.SendGrid.Active {
		return nil
	}

	message := mail.NewSingleEmail(mail.NewEmail("Pandora", "pandora@noreply.com"), subject, mail.NewEmail("", to), "plaint text doesn't supported", html)
	client := sendgrid.NewSendClient(conf.Conf.SendGrid.Token)
	res, err := client.Send(message)
	if err != nil {
		return errors.WithStack(err)
	}
	if res.StatusCode != http.StatusAccepted {
		log.Error(errors.New(fmt.Sprintf("mail doesn't send. http status: %d", res.StatusCode)))
		return errors.ErrMail
	}

	return nil
}
