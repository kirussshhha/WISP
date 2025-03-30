package interfaces

import (
	"WISP/internal/core/domain"
	"github.com/google/uuid"
)

type TeamServiceInterface interface {
	CreateTeam(team *domain.Team) (*domain.Team, error)
	GetTeams() ([]*domain.Team, error)
	GetTeamByID(id uuid.UUID) (*domain.Team, error)
	UpdateTeam(team *domain.Team) (*domain.Team, error)
	DeleteTeam(id uuid.UUID) error
}