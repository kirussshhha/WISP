package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	return r
}

func RegisterHandlers(lifecycle fx.Lifecycle, r *gin.Engine, timeHandler *TimeHandler) {
	timeHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info().Msg("Запуск HTTP сервера на :8080")
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal().Err(err).Msg("Ошибка запуска сервера")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Остановка HTTP сервера")
			return server.Shutdown(ctx)
		},
	})
}
