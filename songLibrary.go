package main

import (
	"embed"
	"log"
	"os"
	"songLibrary/internal/apiserver"
	"songLibrary/internal/config"
	"songLibrary/internal/infrastructure/outsideapi"
	"songLibrary/internal/infrastructure/postgresql"
	"songLibrary/internal/logger/logrus"
	"songLibrary/internal/service"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	//config
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	//logger
	file, err := os.OpenFile(config.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger, err := logrus.New(config.LogLevel, os.Stdout, file)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()
	//store
	store, err := postgresql.New(config.DatabaseConnectString, logger, embedMigrations)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()
	//outsideApi
	outsideApi := outsideapi.New(config.OutsideServerBindAddress, logger)
	//service
	service := service.New(store, outsideApi, logger)
	//apiserver
	logger.Infof("API Server 'Song Library' is started in addr:[%s]", config.BindAddress)
	apiServer := apiserver.New(config.BindAddress, service, logger)
	if err := apiServer.Run(); err != nil {
		logger.Errorf("API Server 'Song Library' error: %s", err)
		return
	}
	logger.Infof("API Server 'Song Library' is stoped")

}
