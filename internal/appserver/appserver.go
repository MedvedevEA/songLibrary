package appserver

import (
	"database/sql"
	"songLibrary/internal/config"
	"songLibrary/internal/controller"
	"songLibrary/internal/infrastructure/logger"
	"songLibrary/internal/infrastructure/outsideapi"
	"songLibrary/internal/infrastructure/postgresql"
	"songLibrary/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Run() error {

	//config
	config, err := config.New()
	if err != nil {
		return err
	}
	//logger
	logger, err := logger.New(config.LogLevel, config.LogFileName)
	if err != nil {
		return err
	}
	defer logger.Close()
	//database
	db, err := sql.Open("postgres", config.DatabaseConnectString)
	if err != nil {
		return err
	}
	defer db.Close()
	//store
	store := postgresql.New(db, logger)
	//outsideApi
	outsideApi := outsideapi.New(config.OutsideServerBindAddress, logger)
	//service
	service := service.New(store, outsideApi, logger)
	//http server
	router := gin.Default()
	//controller
	controller.Init(service, router, logger)
	if err := router.Run(config.BindAddress); err != nil {
		return err
	}

	return nil

}
