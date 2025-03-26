package dbm

import (
	"WISP/internal/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name        string    `gorm:"type:varchar;not null"`
	Description string    `gorm:"type:varchar;not null"`
	CreatedAt   time.Time `gorm:"type:timestamptz"`
	UpdatedAt   time.Time `gorm:"type:timestamptz"`
}

func (t Team) To() *domain.Team {
	return &domain.Team{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func NewTeamDBM(dto *domain.Team) *Team {
	return &Team{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}