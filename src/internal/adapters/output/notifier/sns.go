package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	log "github.com/sirupsen/logrus"
)

type Notifier struct {
	log    *log.Logger
	client *sns.SNS
}

func NewNotifier(l *log.Logger) *Notifier {
	snss := session.Must(session.NewSession())
	client := sns.New(snss)
	return &Notifier{
		log:    l,
		client: client,
	}
}

func (n Notifier) Send_Messages(subject, telephone string) error {
	params := &sns.PublishInput{
		Message:     aws.String(subject),
		PhoneNumber: aws.String(telephone),
	}

	result, err := n.client.Publish(params)
	if err != nil {
		n.log.Errorf("Error sending SMS: %s", err)
		return err
	}

	n.log.Infof("SMS sent. Message ID: %s", *result.MessageId)
	return nil
}
