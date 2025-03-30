package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)

func (s *Services) CreateTeam(team *domain.Team) (*domain.Team, error) {
	team.ID = uuid.New()
	res, err := s.r.DB.CreateTeam(team)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetTeamByID(id uuid.UUID) (*domain.Team, error) {
	res, err := s.r.DB.GetTeamByID(id)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetTeams() ([]*domain.Team, error) {
	res, err := s.r.DB.GetTeams()
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) UpdateTeam(team *domain.Team) (*domain.Team, error) {
	res, err := s.r.DB.UpdateTeam(team)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) DeleteTeam(id uuid.UUID) error {
	err := s.r.DB.DeleteTeam(id)
	if err != nil {
		return domain.ErrInternal
	}

	return nil
}
