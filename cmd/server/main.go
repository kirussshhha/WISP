package main

import (
	"WISP/internal/adapters/handlers"
	"WISP/internal/core/service"
	"WISP/internal/pkg/logger"
	"context"
	"log"
	"time"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			service.NewTimeService,
			handlers.NewTimeHandler,
			handlers.NewGinEngine,
		),
		fx.Invoke(handlers.RegisterHandlers),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		log.Fatalf("Ошибка запуска приложения: %v", err)
	}

	<-app.Done()

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Fatalf("Ошибка остановки приложения: %v", err)
	}
}
