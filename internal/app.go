package app

import (
	"WISP/internal/adapters/graphql"
	"WISP/internal/adapters/handlers"
	"WISP/internal/core/service"
	"WISP/internal/pkg/logger"

	"go.uber.org/fx"
)

var RootApp = fx.New(
	service.Module,
	handlers.Module,
	logger.Module,
	graphql.Module,
)
