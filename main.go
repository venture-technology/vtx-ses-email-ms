package main

import (
	"log"

	"github.com/venture-technology/vtx-ses-emails-ms/config"
	"github.com/venture-technology/vtx-ses-emails-ms/internal/consumer"
	"github.com/venture-technology/vtx-ses-emails-ms/internal/repository"
	"github.com/venture-technology/vtx-ses-emails-ms/internal/service"
)

func main() {

	config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	repo, err := repository.NewPostgres()
	if err != nil {
		log.Fatalf("error creating repository: %s", err.Error())
	}

	aws, err := repository.NewAwsConnection()
	if err != nil {
		log.Fatalf("error creating aws connection: %s", err.Error())
	}

	service := service.New(repo, aws)

	consumer := consumer.New(service)

	log.Printf("initing service: %s", config.Name)
	consumer.Start()

}
