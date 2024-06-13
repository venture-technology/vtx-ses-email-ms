package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/gurodrigues-dev/venture-microservice-emails/config"
	"github.com/gurodrigues-dev/venture-microservice-emails/types"
)

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() (*Postgres, error) {

	conf := config.Get()

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &Postgres{
		conn: db,
	}

	err = repo.migrate(conf.Database.Schema)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (p *Postgres) migrate(filepath string) error {

	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = p.conn.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CreateRecordOfEmailSend(ctx context.Context, email *types.Email) error {

	sqlQuery := `INSERT INTO email_records (recipient, subject, body) VALUES ($1, $2, $3)`
	_, err := p.conn.Exec(sqlQuery, email.Recipient, email.Subject, email.Body)
	return err

}
