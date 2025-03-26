package resolvers

import (
    "WISP/internal/adapters/graphql/generated"
    "WISP/internal/adapters/graphql/model"
    "WISP/internal/core/service"
    "context"
)

type Resolver struct {
    services service.ServicesInterface
}

func NewResolver(services service.ServicesInterface) *Resolver {
    return &Resolver{
        services: services,
    }
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
    return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
    return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
    notes, err := r.services.ListNotes()
    if err != nil {
        return nil, err
    }

    var result []*model.Note
    for _, note := range notes {
        result = append(result, &model.Note{
            ID:      note.ID,
            Title:   note.Title,
            Content: note.Content,
        })
    }
    return result, nil
}

func (r *queryResolver) Note(ctx context.Context, id string) (*model.Note, error) {
    note, err := r.services.GetNoteByID(id)
    if err != nil {
        return nil, err
    }

    return &model.Note{
        ID:      note.ID,
        Title:   note.Title,
        Content: note.Content,
    }, nil
}

func (r *mutationResolver) CreateNote(ctx context.Context, input model.CreateNoteInput) (*model.Note, error) {
    note, err := r.services.CreateNote(input.Title, input.Content)
    if err != nil {
        return nil, err
    }

    return &model.Note{
        ID:      note.ID,
        Title:   note.Title,
        Content: note.Content,
    }, nil
}