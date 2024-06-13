package repository

import (
	"context"

	"github.com/gurodrigues-dev/venture-microservice-emails/types"
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateRecordOfEmailSend(ctx context.Context, email *types.Email) error
}

type Cloud interface {
	SendEmail(ctx context.Context, email *types.Email) error
}
