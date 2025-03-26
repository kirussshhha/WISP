package service

import (
	"WISP/internal/adapters/repositories"
	"WISP/internal/interfaces"

	"go.uber.org/fx"
)

type Services struct {
	r *repositories.Repository
}

type ServicesInterface interface {
	interfaces.UserServiceInterface
	interfaces.TeamMembersServiceInterface
	interfaces.TeamServiceInterface
	interfaces.NoteServiceInterface
	interfaces.TimeServiceInterface
}

var Module = fx.Options(
	fx.Provide(
		NewServices,
	),
)

func NewServices(r *repositories.Repository) ServicesInterface {
	return &Services{
		r,
	}
}
