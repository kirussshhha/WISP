package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)


func (s *Services) CreateTeamMember(userID uuid.UUID, teamID uuid.UUID) (*domain.TeamMember, error) {
	return s.r.DB.CreateTeamMember(userID, teamID)
}

func (s *Services) GetTeamMembers() ([]*domain.TeamMember, error) {
	return s.r.DB.GetTeamMembers()
}

func (s *Services) RemoveTeamMember(userID uuid.UUID, teamID uuid.UUID) error {
	return s.r.DB.RemoveTeamMember(userID, teamID)
}