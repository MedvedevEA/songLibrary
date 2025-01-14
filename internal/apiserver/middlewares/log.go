package middlewares

import (
	"songLibrary/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		latencyTime := time.Since(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		logger.Debugf("| HTTP request | %d | %v | %s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)

		ctx.Next()
	}
}
