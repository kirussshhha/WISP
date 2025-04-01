package repositories

import (
	"WISP/internal/adapters/repositories/database"
	"WISP/internal/config"
	"WISP/internal/interfaces"

	"go.uber.org/fx"
)

type Repository struct {
	DB   database.Repository
	Note interfaces.NoteRepository
}

var Module = fx.Options(
	fx.Provide(
		New,
		database.NewNoteRepository,
	),
)

func New(config *config.DBConfig, noteRepo *database.NoteRepository) (*Repository, error) {
	db, err := database.New(config)
	if err != nil {
		return nil, err
	}

	return &Repository{
		DB: db,
		Note: noteRepo,
	}, nil
}
