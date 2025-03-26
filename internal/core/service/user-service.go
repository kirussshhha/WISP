package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)

func (s *Services) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = uuid.New()
    return s.r.DB.CreateUser(user)
}

func (s *Services) GetUserByEmail(email string) (*domain.User, error) {
    return s.r.DB.GetUserByEmail(email)
}

func (s *Services) GetUserByID(id uuid.UUID) (*domain.User, error) {
	return s.r.DB.GetUserByID(id)
}

func (s *Services) DeleteUser(id uuid.UUID) error {
	return s.r.DB.DeleteUser(id)
}