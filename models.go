package main

// note data structure

type Note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DB :)
var notes = []*Note{
	{ID: "1", Title: "Sneed's Feed and Seed", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sit amet volutpat lacus. Cras et."},
	{ID: "2", Title: "With regards to a skibidi toilet", Content: "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit..."},
	{ID: "3", Title: "Formerly Chuck's", Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lorem urna, sollicitudin ut fermentum eu, efficitur nec quam. Duis a."},
}

func listNotes() []*Note {
	return notes
}

func getNote(id string) *Note {
	for _, note := range notes {
		if note.ID == id {
			return note
		}
	}
	return nil
}

func addNote(note Note) {
	notes = append(notes, &note)
}
