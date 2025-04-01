package handlers

import (
	"WISP/internal/adapters/graphql"
	"WISP/internal/adapters/handlers/middleware"
	"WISP/internal/core/service"
	"WISP/internal/pkg/rabbitmq"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHTTPServer),
	fx.Invoke(registerRoutes),
	fx.Invoke(runHTTPServer),
)

type Handler struct {
	Gin            *gin.Engine
	Services       service.ServicesInterface
	GraphQLHandler *graphql.GraphQLHandler
	Producer       *rabbitmq.Producer
}

func NewHTTPServer(
	service service.ServicesInterface,
	gqHandler *graphql.GraphQLHandler,
	producer *rabbitmq.Producer,
) *Handler {
	engine := gin.Default()

	engine.Use(
		middleware.NewLoggerMiddleware(),
	)

	return &Handler{
		Gin:            engine,
		Services:       service,
		GraphQLHandler: gqHandler,
		Producer:       producer,
	}
}

func registerRoutes(h *Handler) {
	v1 := h.Gin.Group("/api/v1")
	{
		v1.GET("/time", h.GetTime)
		v1.GET("/time/:format", h.GetTimeWithPathFormat)
		v1.GET("/time-query", h.GetTimeWithQueryFormat)
		v1.POST("/time-diff", h.CalculateTimeDifference)

		v1.POST("/test-emails", h.SendTestEmails)

		user := v1.Group("/user")
		{
			user.POST("", h.CreateUser)
			user.GET("", h.GetUsers)
			user.GET("/:id", h.GetUserByID)
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)
		}
		team := v1.Group("/team")
		{
			team.POST("", h.CreateTeam)
			team.GET("", h.GetTeams)
			team.GET("/:id", h.GetTeamByID)
			team.PUT("/:id", h.UpdateTeam)
			team.DELETE("/:id", h.DeleteTeam)
		}
		teamMember := v1.Group("/team-member")
		{
			teamMember.GET("", h.GetTeamMembers)
			teamMember.POST("/:userId/invite/:teamId", h.CreateTeamMember)
			teamMember.DELETE("/:userId/leave/:teamId", h.RemoveTeamMember)
		}
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
			log.Info().Msg("Starting Application")
			go h.Gin.Run(":8080")
			return nil
		},
		OnStop: func(context.Context) error {
			log.Info().Msg("Stopping Application")
			return nil
		},
	})
}
