package graphql

import (
	"WISP/internal/adapters/graphql/generated"
    "WISP/internal/adapters/graphql/resolvers"
    "net/http"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gin-gonic/gin"
)

type GraphQLHandler struct {
    Schema     *handler.Server
    Playground http.HandlerFunc
}

func NewGraphQLHandler(resolver *resolvers.Resolver) *GraphQLHandler {
	config := generated.Config{
		Resolvers: resolver,
	}

    return &GraphQLHandler{
        Schema:     handler.NewDefaultServer(generated.NewExecutableSchema(config)), 
        Playground: playground.Handler("GraphQL", "/query"),
    }
}

func (h *GraphQLHandler) RegisterRoutes(r *gin.Engine) {
    r.POST("/query", gin.WrapH(h.Schema))
    r.GET("/playground", gin.WrapH(h.Playground))
}