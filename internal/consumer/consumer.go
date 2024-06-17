package consumer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
	"github.com/venture-technology/vtx-ses-emails-ms/config"
	"github.com/venture-technology/vtx-ses-emails-ms/internal/service"
)

type consumer struct {
	service *service.Service
}

func New(s *service.Service) *consumer {
	return &consumer{
		service: s,
	}
}

func (ct *consumer) Start() {

	conf := config.Get()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{conf.Messaging.Brokers},
		Topic:     conf.Messaging.Topic,
		Partition: 1,
		GroupID:   "reader.kafka.group",
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer reader.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {

			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("Erro ao ler mensagem do Kafka: %v", err)
			}

			email, err := ct.service.UnserializeJsonToEmailDto(context.Background(), &message)
			if err != nil {
				log.Fatalf("Erro ao unserializar mensagem do Kafka: %v", err)
			}
			log.Printf("Message to -->: %s", email)

			err = ct.service.SendEmail(context.Background(), email)
			if err != nil {
				log.Fatalf("Erro ao enviar email: %v", err)
			}

			err = ct.service.CreateRecordOfEmailSend(context.Background(), email)
			if err != nil {
				log.Fatalf("Erro ao gravar record do email: %v", err)
			}

			log.Println("Venture-Microservice-Email: Message found in Queue, Email sended.")
		}
	}()

	<-signals

}
