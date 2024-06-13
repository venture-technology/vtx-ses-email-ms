package repository

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gurodrigues-dev/venture-microservice-emails/config"
	"github.com/gurodrigues-dev/venture-microservice-emails/types"
)

type AWS struct {
	conn *session.Session
}

func NewAwsConnection() (*AWS, error) {

	conf := config.Get()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Cloud.Region),
		Credentials: credentials.NewStaticCredentials(conf.Cloud.AccessKey, conf.Cloud.SecretKey, conf.Cloud.Token),
	})

	if err != nil {
		return nil, err
	}

	repo := &AWS{
		conn: sess,
	}

	return repo, nil

}

func (a *AWS) SendEmail(ctx context.Context, email *types.Email) error {

	conf := config.Get()

	svc := ses.New(a.conn)

	emailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(email.Recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(email.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(email.Subject),
			},
		},
		Source: aws.String(conf.Cloud.Source),
	}

	_, err := svc.SendEmail(emailInput)

	if err != nil {
		return err
	}

	return nil

}
