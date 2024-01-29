package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//health check
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("WORKING FINE(I GUESS)"))
	})

	r.Mount("/v1/note", NoteRoutes())

	http.ListenAndServe(":8080", r)
}

func NoteRoutes() chi.Router {
	r := chi.NewRouter()
	noteHandler := NoteHandler{}
	r.Get("/", noteHandler.ListNotes)
	r.Post("/", noteHandler.CreateNote)
	r.Get("/{id}", noteHandler.GetNote)
	return r
}
