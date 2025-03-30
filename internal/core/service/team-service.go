package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (s *Services) CreateTeam(team *domain.Team) (*domain.Team, error) {
	team.ID = uuid.New()
	res, err := s.r.DB.CreateTeam(team)
	if err != nil {
		log.Error().Err(err).Str("service", "createTeam").Msg("Failed to create team")
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetTeamByID(id uuid.UUID) (*domain.Team, error) {
	res, err := s.r.DB.GetTeamByID(id)
	if err != nil {
		log.Error().Err(err).Str("service", "getTeamByID").Msg("Failed to get team by ID")
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetTeams() ([]*domain.Team, error) {
	res, err := s.r.DB.GetTeams()
	if err != nil {
		log.Error().Err(err).Str("service", "getTeams").Msg("Failed to get teams")
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) UpdateTeam(team *domain.Team) (*domain.Team, error) {
	res, err := s.r.DB.UpdateTeam(team)
	if err != nil {
		log.Error().Err(err).Str("service", "updateTeam").Msg("Failed to update team")
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) DeleteTeam(id uuid.UUID) error {
	err := s.r.DB.DeleteTeam(id)
	if err != nil {
		log.Error().Err(err).Str("service", "deleteTeam").Msg("Failed to delete team")
		return domain.ErrInternal
	}

	return nil
}
