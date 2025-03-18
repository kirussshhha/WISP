package graphql

import (
	"WISP/internal/adapters/graphql/resolvers"

	"go.uber.org/fx"
)

var Module = fx.Options(
    fx.Provide(
		resolvers.NewResolver,
		NewGraphQLHandler,
	),
)