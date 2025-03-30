package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (s *Services) CreateTeamMember(userID uuid.UUID, teamID uuid.UUID) (*domain.TeamMember, error) {
	res, err := s.r.DB.CreateTeamMember(userID, teamID)
	if err != nil {
		log.Error().Err(err).Str("service", "createTeamMember").Msg("Failed to create team member")
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetTeamMembers() ([]*domain.TeamMember, error) {
	res, err := s.r.DB.GetTeamMembers()
	if err != nil {
		log.Error().Err(err).Str("service", "getTeamMembers").Msg("Failed to get team members")
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) RemoveTeamMember(userID uuid.UUID, teamID uuid.UUID) error {
	err := s.r.DB.RemoveTeamMember(userID, teamID)
	if err != nil {
		log.Error().Err(err).Str("service", "removeTeamMember").Msg("Failed to remove team member")
		return domain.ErrInternal
	}

	return nil
}
