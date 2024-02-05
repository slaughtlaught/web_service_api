package entity

import "github.com/slaughtlaught/web-service-api/internal/value"

type User struct {
	ID             value.NoteID `json:"id"`
	Name           string       `json:"name"`
	Email          value.Email  `json:"email"`
	HashedPassword string       `json:"hashedPassword"`
}
