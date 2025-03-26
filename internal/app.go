package app

import (
	"WISP/internal/adapters/graphql"
	"WISP/internal/adapters/handlers"
	"WISP/internal/adapters/repositories"
	"WISP/internal/core/service"
	"WISP/internal/pkg/logger"

	"go.uber.org/fx"
)

var RootApp = fx.New(
	repositories.Module,
	service.Module,
	handlers.Module,
	logger.Module,
	graphql.Module,
)
