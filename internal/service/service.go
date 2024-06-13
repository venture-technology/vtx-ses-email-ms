package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gurodrigues-dev/venture-microservice-emails/types"
	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/vtx-ses-emails-ms/internal/repository"
)

type Service struct {
	repository repository.Repository
	cloud      repository.Cloud
}

func New(repo repository.Repository, cloud repository.Cloud) *Service {
	return &Service{
		repository: repo,
		cloud:      cloud,
	}
}

func (s *Service) CreateRecordOfEmailSend(ctx context.Context, recipient *types.Email) error {
	return s.repository.CreateRecordOfEmailSend(ctx, recipient)
}

func (s *Service) SendEmail(ctx context.Context, email *types.Email) error {
	return s.cloud.SendEmail(ctx, email)
}

func (s *Service) UnserializeJsonToEmailDto(ctx context.Context, msg *kafka.Message) (*types.Email, error) {
	var email *types.Email

	err := json.Unmarshal(msg.Value, &email)
	if err != nil {
		log.Fatalf("Erro ao desserializar mensagem JSON: %v", err)
		return nil, err
	}

	return email, nil
}
