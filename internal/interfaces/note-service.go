package interfaces

import "WISP/internal/core/domain"

type NoteServiceInterface interface {
	CreateNote(title, content string) (*domain.Note, error)
	GetNoteByID(id string) (*domain.Note, error)
    ListNotes() ([]*domain.Note, error)
}