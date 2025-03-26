package interfaces

import (
	"WISP/internal/core/domain"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
    CreateUser(user *domain.User) (*domain.User, error)
    GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uuid.UUID) (*domain.User, error)	
	DeleteUser(id uuid.UUID) error
}