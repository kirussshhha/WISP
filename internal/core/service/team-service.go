package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)

func (s *Services) CreateTeam(team *domain.Team) (*domain.Team, error) {
	team.ID = uuid.New()
	return s.r.DB.CreateTeam(team)
}

func (s *Services) GetTeams() ([]*domain.Team, error) {
	return s.r.DB.GetTeams()
}

func (s *Services) UpdateTeam(team *domain.Team) (*domain.Team, error) {
	return s.r.DB.UpdateTeam(team)
}

func (s *Services) DeleteTeam(id uuid.UUID) error {
	return s.r.DB.DeleteTeam(id)
}