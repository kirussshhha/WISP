package app

import (
	"WISP/internal/adapters/graphql"
	"WISP/internal/adapters/handlers"
	"WISP/internal/adapters/handlers/middleware"
	"WISP/internal/adapters/repositories"
	"WISP/internal/core/service"
	"WISP/internal/pkg/logger"

	"go.uber.org/fx"
)

var RootApp = fx.New(
	logger.Module,
	repositories.Module,
	service.Module,
	handlers.Module,
	graphql.Module,
	middleware.Module,
)
