package dbm

import (
	"WISP/internal/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"type:timestamptz"`
	UpdatedAt time.Time `gorm:"type:timestamptz"`
	Username  string    `gorm:"type:varchar;not null"`
	Email     string    `gorm:"type:varchar;not null"`
	Password  string    `gorm:"type:varchar;not null"`
}

func (u User) To() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewUserDBM(dto *domain.User) *User {
	return &User{
		ID:       dto.ID,
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
