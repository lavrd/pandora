package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"pandora/pkg/conf"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
)

// Mail
type Mail struct {
	Personalizations []*personalizations `json:"personalizations"`
	Subject          string              `json:"subject"`
	From             *email              `json:"from"`
	Content          []*content          `json:"content"`
}

type personalizations struct {
	To []*email `json:"to"`
}

type content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type email struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func send(to, subject, html string) error {
	if !conf.Conf.SendGrid.Active {
		return nil
	}

	mail := &Mail{
		Personalizations: []*personalizations{
			{
				To: []*email{
					{
						Email: to,
					},
				},
			},
		},
		Subject: subject,
		From: &email{
			Name:  "Pandora",
			Email: "pandora@mail.sys",
		},
		Content: []*content{
			{
				Type:  "text/html",
				Value: html,
			},
		},
	}

	buf, err := json.Marshal(mail)
	if err != nil {
		return errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.sendgrid.com/v3/mail/send", bytes.NewBuffer(buf))
	if err != nil {
		return errors.WithStack(err)
	}
	defer req.Body.Close()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", conf.Conf.SendGrid.Token))
	req.Header.Add("Content-Type", "application/json")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		log.Error(errors.New(fmt.Sprintf("mail doesn't send. http status: %d", res.StatusCode)))
		return errors.ErrMail
	}

	return nil
}
