package service

import (
	"WISP/internal/core/domain"
	"fmt"

	"github.com/google/uuid"
)

func (s *Services) CreateNote(title, content string) (*domain.Note, error) {
	note := &domain.Note{
		ID:      uuid.New().String(),
		Title:   title,
		Content: content,
	}
	s.notes[note.ID] = note
	return note, nil
}

func (s *Services) GetNoteByID(id string) (*domain.Note, error) {
	if note, exist := s.notes[id]; exist {
		return note, nil
	}
	return nil, fmt.Errorf("note not found")
}

func (s *Services) ListNotes() ([]*domain.Note, error) {
	notes := make([]*domain.Note, 0, len(s.notes))
	for _, note := range s.notes {
		notes = append(notes, note)
	}
	return notes, nil
}