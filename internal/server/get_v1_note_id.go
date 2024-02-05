package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/slaughtlaught/web-service-api/internal/value"
)

func (n NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	parsedID, err := value.ParseNoteID(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	note, err := n.noteService.GetByID(r.Context(), parsedID)
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
