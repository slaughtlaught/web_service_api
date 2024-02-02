package main

// note data structure

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

//func listNotes() []*Note {
//	return notes
//}

//func getNote(id string) *Note {
//	for _, note := range notes {
//		if note.ID == id {
//			return note
//		}
//	}
//	return nil
//}

//func addNote(note Note) {
//	notes = append(notes, &note)
//}
