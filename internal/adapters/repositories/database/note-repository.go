package database

import (
    "WISP/internal/core/domain"
    "fmt"
)

type NoteRepository struct {
    notes map[string]*domain.Note
}

func NewNoteRepository() *NoteRepository {
    return &NoteRepository{
        notes: make(map[string]*domain.Note),
    }
}

func (r *NoteRepository) Create(note *domain.Note) (*domain.Note, error) {
    r.notes[note.ID] = note
    return note, nil
}

func (r *NoteRepository) GetByID(id string) (*domain.Note, error) {
    if note, exist := r.notes[id]; exist {
        return note, nil
    }
    return nil, fmt.Errorf("note not found")
}

func (r *NoteRepository) List() ([]*domain.Note, error) {
    notes := make([]*domain.Note, 0, len(r.notes))
    for _, note := range r.notes {
        notes = append(notes, note)
    }
    return notes, nil
}