package server

import (
	"encoding/json"
	"net/http"

	"github.com/slaughtlaught/web-service-api/internal/entity"
)

func (n NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note entity.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if err := n.noteService.Add(r.Context(), note); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
