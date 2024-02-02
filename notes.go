package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteHandler struct {
	db *pgxpool.Pool
}

func (n NoteHandler) ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(GetAllNotes())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (n NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	note := GetNoteById(id)
	err := json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (n NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AddNote(note)
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
