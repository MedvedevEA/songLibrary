package apiserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"songLibrary/internal/controller"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	server *http.Server
}

func New(bindAddress string, service controller.Service, logger controller.Logger) *ApiServer {
	router := gin.New()
	controller.Init(router, service, logger)

	server := &http.Server{
		Addr:    bindAddress,
		Handler: router,
	}
	return &ApiServer{
		server: server,
	}
}
func (a *ApiServer) Run() error {
	chError := make(chan error, 1)
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			chError <- err
		}
	}()
	go func() {
		chQuit := make(chan os.Signal, 1)
		signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)
		<-chQuit
		ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)
		defer channel()
		chError <- a.server.Shutdown(ctx)
	}()

	return <-chError
}
