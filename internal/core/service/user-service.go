package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)

func (s *Services) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = uuid.New()
	res, err := s.r.DB.CreateUser(user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetUsers() ([]*domain.User, error) {
	res, err := s.r.DB.GetUsers()
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetUserByEmail(email string) (*domain.User, error) {
	res, err := s.r.DB.GetUserByEmail(email)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) GetUserByID(id uuid.UUID) (*domain.User, error) {
	res, err := s.r.DB.GetUserByID(id)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) UpdateUser(user *domain.User) (*domain.User, error) {
	res, err := s.r.DB.UpdateUser(user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return res, nil
}

func (s *Services) DeleteUser(id uuid.UUID) error {
	err := s.r.DB.DeleteUser(id)
	if err != nil {
		return domain.ErrInternal
	}

	return nil
}
