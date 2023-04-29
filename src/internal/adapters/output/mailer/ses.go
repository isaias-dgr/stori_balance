package adapter

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/isaias-dgr/story-balance/src/internal/core/domain"
	log "github.com/sirupsen/logrus"
)

type Mailer struct {
	log      *log.Logger
	client   *ses.SES
	source   string
	template string
}

func NewMailer(s, t string, l *log.Logger) *Mailer {
	sess := session.Must(session.NewSession())
	client := ses.New(sess)
	return &Mailer{
		log:      l,
		client:   client,
		source:   s,
		template: t,
	}
}

func (m Mailer) Send(dest string, acc domain.Account) error {
	m.log.Debugf("%s %s %s", dest, m.source, m.template)
	jsonData, err := json.Marshal(acc)
	if err != nil {
		m.log.Errorf("Error al convertir a JSON: %s", err)
		return err
	}

	input := &ses.SendTemplatedEmailInput{
		Source: aws.String(m.source),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(dest),
			},
		},
		Template:     aws.String(m.template),
		TemplateData: aws.String(string(jsonData)),
	}

	result, err := m.client.SendTemplatedEmail(input)
	if err != nil {
		m.log.Errorf("Error sending email: %s", err.Error())
		return err
	}

	m.log.Infof("Email sent successfully %s", *result.MessageId)
	return nil
}
