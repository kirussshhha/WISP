package app

import (
	"WISP/internal/adapters/graphql"
	"WISP/internal/adapters/handlers"
	"WISP/internal/adapters/handlers/middleware"
	"WISP/internal/adapters/repositories"
	"WISP/internal/config"
	"WISP/internal/core/service"
	"WISP/internal/pkg/logger"
	"WISP/internal/pkg/rabbitmq"

	"go.uber.org/fx"
)

var RootApp = fx.New(
	config.Module,
	logger.Module,
	rabbitmq.Module, 
	repositories.Module,
	service.Module,
	handlers.Module,
	graphql.Module,
	middleware.Module,
)
