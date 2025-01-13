package controller

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
}
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}
type Controller struct {
	Service Service
	Logger  Logger
}

func Init(router *gin.Engine, service Service, logger Logger) {
	controller := &Controller{
		Service: service,
		Logger:  logger,
	}
	router.GET("test", controller.Test)
}
func (c *Controller) Test(ctx *gin.Context) {
	c.Logger.Debugf("Test route", nil)
	ctx.Status(200)
}
