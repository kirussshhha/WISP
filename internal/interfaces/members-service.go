package interfaces

import (
	"WISP/internal/core/domain"
	"github.com/google/uuid"
)


type TeamMembersServiceInterface interface {
	CreateTeamMember(userID uuid.UUID, teamID uuid.UUID) (*domain.TeamMember, error)
	GetTeamMembers() ([]*domain.TeamMember, error)
	RemoveTeamMember(userID uuid.UUID, teamID uuid.UUID) error
}