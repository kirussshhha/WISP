package main

import (
	"os"
	"time"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Msg("HTTP request") 
	}
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	r := gin.New()
	r.Use(Logger()) 

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	r.Run(":8080")
}
