package server

import (
	"encoding/json"
	"net/http"
)

func (n NoteHandler) ListNotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	notes, err := n.noteService.List(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
