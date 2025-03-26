package service

import (
	"WISP/internal/core/domain"

	"github.com/google/uuid"
)

func (s *Services) CreateNote(title, content string) (*domain.Note, error) {
	note := &domain.Note{
		ID:      uuid.New().String(),
		Title:   title,
		Content: content,
	}
	return s.r.Note.Create(note)
}

func (s *Services) GetNoteByID(id string) (*domain.Note, error) {
	return s.r.Note.GetByID(id)
}

func (s *Services) ListNotes() ([]*domain.Note, error) {
	return s.r.Note.List()
}
