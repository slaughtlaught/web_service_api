package value

import (
	"fmt"

	"github.com/rs/xid"
)

type NoteID struct {
	xid.ID
}

func ParseNoteID(s string) (NoteID, error) {
	id, err := xid.FromString(s)
	if err != nil {
		return NoteID{}, fmt.Errorf("xid.FromString: %w", err)
	}
	return NoteID{ID: id}, nil
}
