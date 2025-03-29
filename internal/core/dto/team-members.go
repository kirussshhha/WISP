package dto

import (
	"time"

	"github.com/google/uuid"
)

type TeamMember struct {
	UserID    uuid.UUID `json:"user_id"`
	TeamID    uuid.UUID `json:"team_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToTeamMemberDTO(
	userID uuid.UUID,
	teamID uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
) *TeamMember {
	return &TeamMember{
		UserID:    userID,
		TeamID:    teamID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
