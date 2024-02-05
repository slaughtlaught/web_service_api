package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/slaughtlaught/web-service-api/internal/entity"
	"github.com/slaughtlaught/web-service-api/internal/value"
)

type noteService interface {
	GetByID(context.Context, value.NoteID) (entity.Note, error)
	List(context.Context) ([]entity.Note, error)
	Add(context.Context, entity.Note) error
}

type NoteHandler struct {
	noteService noteService
}

func NewNoteHandler(noteService noteService) NoteHandler {
	return NoteHandler{noteService: noteService}
}

func NewRouter(n NoteHandler) chi.Router {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Second * 5)) //health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("WORKING FINE(I GUESS)"))
	})
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * 4):
		}
		w.Write([]byte("done"))

	})

	r.Route("/v1/note", func(r chi.Router) {
		r.Get("/", n.ListNotes)
		r.Post("/", n.CreateNote)
		r.Get("/{id}", n.GetNote)
	})

	return r
}
