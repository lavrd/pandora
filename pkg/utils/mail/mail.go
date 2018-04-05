package mail

import (
	"bytes"
	"encoding/json"
	`errors`
	`fmt`
	"net/http"

	`github.com/spacelavr/pandora/pkg/log`
	`github.com/spf13/viper`
)

type mail struct {
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

// Send send mail to email address
func Send(to, subject, html string) error {

	log.Debugf("send mail to %s", to)

	var mail = &mail{
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
			Name:  viper.GetString("sendgrid.name"),
			Email: viper.GetString("sendgrid.email"),
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

	req, err := http.NewRequest("POST", viper.GetString("sendgrid.endpoint"), bytes.NewBuffer(buf))
	if err != nil {
		log.Error(err)
		return err
	}
	defer req.Body.Close()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("sendgrid.token")))
	req.Header.Add("Content-Type", "application/json")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		log.Errorf("mail doesn't send. http status: %d", res.StatusCode)
		return errors.New("send mail error")
	}

	return nil
}
