package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, closefunc, err := NewStore()
	if err != nil {
		log.Printf("error connecting to a database %v", err)
	}
	defer closefunc()
	if err1 := CreateTable(pool); err1 != nil {
		log.Printf("error creating a table database %v", err1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("WORKING FINE(I GUESS)"))
	})

	r.Mount("/v1/note", NoteRoutes(pool))

	http.ListenAndServe(":8080", r)
}

func NoteRoutes(pool *pgxpool.Pool) chi.Router {
	r := chi.NewRouter()
	noteHandler := NoteHandler{db: pool}
	r.Get("/", noteHandler.ListNotes)
	r.Post("/", noteHandler.CreateNote)
	r.Get("/{id}", noteHandler.GetNote)
	return r
}
