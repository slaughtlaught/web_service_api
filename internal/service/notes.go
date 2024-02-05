package service

import (
	"context"
	"fmt"

	"github.com/slaughtlaught/web-service-api/internal/entity"
	"github.com/slaughtlaught/web-service-api/internal/value"
)

type noteSource interface {
	GetNoteByID(context.Context, value.NoteID) (entity.Note, error)
	GetAllNotes(context.Context) ([]entity.Note, error)
	AddNote(context.Context, entity.Note) error
}

type Notes struct {
	noteSource noteSource
}

func NewNotes(noteSource noteSource) Notes {
	return Notes{
		noteSource: noteSource,
	}
}

func (n Notes) List(ctx context.Context) ([]entity.Note, error) {

	notes, err := n.noteSource.GetAllNotes(ctx)
	if err != nil {
		return nil, fmt.Errorf("noteSource.GetAllNotes: %w", err)
	}

	return notes, nil
}

func (n Notes) GetByID(ctx context.Context, id value.NoteID) (entity.Note, error) {
	note, err := n.noteSource.GetNoteByID(ctx, id)
	if err != nil {
		return entity.Note{}, fmt.Errorf("noteSource.GetNoteByID: %w", err)
	}

	return note, nil
}

func (n Notes) Add(ctx context.Context, note entity.Note) error {
	if err := n.noteSource.AddNote(ctx, note); err != nil {
		return fmt.Errorf("noteSource.AddNote: %w", err)
	}
	return nil
}
