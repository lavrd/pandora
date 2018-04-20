package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spf13/viper"
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
	if !viper.GetBool("mail.send") {
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
			Name:  viper.GetString("mail.name"),
			Email: viper.GetString("mail.email"),
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
		log.Error(err)
		return err
	}

	req, err := http.NewRequest("POST", viper.GetString("mail.endpoint"), bytes.NewBuffer(buf))
	if err != nil {
		log.Error(err)
		return err
	}
	defer req.Body.Close()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("mail.token")))
	req.Header.Add("Content-Type", "application/json")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		log.Errorf("mail doesn't send. http status: %d", res.StatusCode)
		return errors.SendMailError
	}

	return nil
}
