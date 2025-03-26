package domain

import (
	"time"

	"github.com/google/uuid"
)

type TeamMember struct {
	UserID    uuid.UUID
	TeamID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
