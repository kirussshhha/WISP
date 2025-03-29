package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)

func (s *Services) CreateTeamMember(userID uuid.UUID, teamID uuid.UUID) (*domain.TeamMember, error) {
	res, err := s.r.DB.CreateTeamMember(userID, teamID)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetTeamMembers() ([]*domain.TeamMember, error) {
	res, err := s.r.DB.GetTeamMembers()
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) RemoveTeamMember(userID uuid.UUID, teamID uuid.UUID) error {
	err := s.r.DB.RemoveTeamMember(userID, teamID)
	if err != nil {
		return domain.ErrInternal
	}

	return nil
}
