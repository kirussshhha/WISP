package handlers

import (
	"WISP/internal/adapters/graphql"
	"WISP/internal/core/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHTTPServer),
	fx.Invoke(registerRoutes),
	fx.Invoke(runHTTPServer),
)

type Handler struct {
	Gin            *gin.Engine
	Services       *service.Services
	GraphQLHandler *graphql.GraphQLHandler
}

func NewHTTPServer(service *service.Services, gqHandler *graphql.GraphQLHandler) *Handler {
	engine := gin.Default()
	return &Handler{
		Gin:            engine,
		Services:       service,
		GraphQLHandler: gqHandler,
	}
}

func registerRoutes(h *Handler) {
	v1 := h.Gin.Group("/api/v1")
	{
		v1.GET("/time", h.GetTime)
		v1.GET("/time/:format", h.GetTimeWithPathFormat)
		v1.GET("/time-query", h.GetTimeWithQueryFormat)
		v1.POST("/time-diff", h.CalculateTimeDifference)
	}

	v2 := h.Gin.Group("/api/v2")
	{
		v2.GET("/time", h.GetTimeV2)
	}

	h.GraphQLHandler.RegisterRoutes(h.Gin)
}

func runHTTPServer(lifecycle fx.Lifecycle, h *Handler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting Application")
			go h.Gin.Run(":8080")
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application")
			return nil
		},
	})
}
