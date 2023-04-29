package adapter

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/sirupsen/logrus"
)

type Queue struct {
	log      *log.Logger
	queueURL string
	client   *sqs.SQS
}

func NewQueue(u string, l *log.Logger) *Queue {
	sess := session.Must(session.NewSession())

	return &Queue{
		log:      l,
		client:   sqs.New(sess),
		queueURL: u,
	}
}

func (s Queue) Send(message *string) error {
	messageAttribute := &sqs.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(aws.ToString(message)),
	}
	result, err := s.client.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(5),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"File": messageAttribute,
		},
		MessageBody: aws.String(aws.ToString(message)),
		QueueUrl:    &s.queueURL,
	})
	if err != nil {
		return err
	}
	s.log.Infof("Message send %s", *result.MessageId)
	return nil
}
