package entity

import "github.com/slaughtlaught/web-service-api/internal/value"

type Note struct {
	ID      value.NoteID `json:"id" db:"id"`
	Title   string       `json:"title" db:"title"`
	Content string       `json:"content" db:"content"`
	UserID  string       `json:"userId" db:"user_id"`
}
