package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type NoteHandler struct {
}

func (n NoteHandler) ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(listNotes())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (n NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	note := getNote(id)
	if note == nil {
		http.Error(w, "Note with given ID not found", http.StatusNotFound)
	}
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
	addNote(note)
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
