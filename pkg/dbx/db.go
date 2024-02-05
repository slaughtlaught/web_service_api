package dbx

import (
	"context"

	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewStorage(ctx context.Context) (*sqlx.DB, func(), error) {
	urlOfDb := "postgres://postgres:postgres@postgres:5432/notes?sslmode=disable"
	db, err := sqlx.Open("pgx", urlOfDb)
	if err != nil {
		return nil, nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	db.SetMaxOpenConns(5)

	if err = db.Ping(); err != nil {
		return nil, nil, fmt.Errorf("db.Ping: %w", err)
	}

	dbClose := func() {
		db.Close()
	}

	return db, dbClose, err

}
