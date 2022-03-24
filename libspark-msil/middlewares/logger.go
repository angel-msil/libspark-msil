package middlewares

import (
	"time"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"libspark-msil/constants"
)

// Logger is the middleware which logs the requests
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path

		// Process request
		ctx.Next()

		// stop timer
		end := time.Now()
		latency := end.Sub(start)

		log.Info().
			Int(constants.StatusCodeKey, ctx.Writer.Status()).
			Dur(constants.LatencyKey, latency).
			Str(constants.ClientIPKey, ctx.ClientIP()).
			Str(constants.MethodKey, ctx.Request.Method).
			Str(constants.PathKey, path).
			Str(constants.ErrorKey, ctx.Errors.ByType(gin.ErrorTypePrivate).String()).
			Int("pid", os.Getpid()).
			Send()
	}
}
