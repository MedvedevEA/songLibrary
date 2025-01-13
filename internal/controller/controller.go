package controller

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
}
type Logger interface {
	Debug(string)
	Debugf(string, []interface{})
	Info(string)
	Infof(string, []interface{})
}
type Controller struct {
	Service Service
	Logger  Logger
}

func Init(service Service, router *gin.Engine, logger Logger) {
	controller := &Controller{
		Service: service,
		Logger:  logger,
	}
	router.GET("test", controller.Test)
}
func (c *Controller) Test(ctx *gin.Context) {
	c.Logger.Debug("Test route")
	ctx.Status(200)
}
