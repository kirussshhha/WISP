package service

import (
	"WISP/internal/core/domain"

	"go.uber.org/fx"
)

type Services struct {
	notes map[string]*domain.Note
}

var Module = fx.Options(
	fx.Provide(
		NewServices,
	),
)

func NewServices() *Services {
	return &Services{
		notes: make(map[string]*domain.Note),
	}
}
