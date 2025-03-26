package dbm

import (
	"WISP/internal/core/domain"
	"time"

	"github.com/google/uuid"
)

type TeamMember struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;autoIncrement:false"`
	TeamID    uuid.UUID `gorm:"type:uuid;primaryKey;autoIncrement:false"`
	CreatedAt time.Time `gorm:"type:timestamptz"`
	UpdatedAt time.Time `gorm:"type:timestamptz"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Team Team `gorm:"foreignKey:TeamID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t TeamMember) To() *domain.TeamMember {
	return &domain.TeamMember{
		UserID:    t.UserID,
		TeamID:    t.TeamID,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func NewTeamMemberDBM(dto *domain.TeamMember) *TeamMember {
	return &TeamMember{
		UserID:    dto.UserID,
		TeamID:    dto.TeamID,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
