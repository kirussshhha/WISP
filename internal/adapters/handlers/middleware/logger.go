package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func NewLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		logger := log.With().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Logger()

		c.Set("logger", logger)

		c.Next()

		logger.Info().
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Msg("request completed")
	}
}
