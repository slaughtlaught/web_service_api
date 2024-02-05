package persistance

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/slaughtlaught/web-service-api/internal/entity"
	"github.com/slaughtlaught/web-service-api/internal/value"
	"github.com/slaughtlaught/web-service-api/pkg/errorx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Storage {
	return Storage{db: db}
}

func (s Storage) GetAllNotes(ctx context.Context) ([]entity.Note, error) {
	notes := []entity.Note{}

	if err := s.db.SelectContext(ctx, &notes, "select * from notes"); err != nil {
		return nil, fmt.Errorf("db.QueryContext: %w", err)
	}

	return notes, nil
}

func (s Storage) GetNoteByID(ctx context.Context, id value.NoteID) (entity.Note, error) {
	n := entity.Note{}

	if err := s.db.GetContext(ctx, &n, "select * from notes where id = $1", id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Note{}, errorx.NewNotFoundError(err)
		}
		return entity.Note{}, fmt.Errorf("db.GetContext: %w", err)
	}

	return n, nil
}

func (s Storage) AddNote(ctx context.Context, note entity.Note) error {
	guid := xid.New()
	id, title, content := guid, note.Title, note.Content

	if _, err := s.db.ExecContext(ctx, "insert into notes (id, title, content) values ($1, $2, $3)", id, title, content); err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}
	return nil
}
