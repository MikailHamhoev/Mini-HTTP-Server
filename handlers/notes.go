// mini-http-server/handlers/notes.go
package handlers

import (
	"mini-http-server/models"
	"sync"
)

type NoteStore struct {
	mu     sync.RWMutex
	notes  map[int]models.Note
	nextID int
}

func NewNoteStore() *NoteStore {
	return &NoteStore{
		notes:  make(map[int]models.Note),
		nextID: 1,
	}
}

func (ns *NoteStore) Create(note models.Note) models.Note {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	note.ID = ns.nextID
	ns.nextID++
	ns.notes[note.ID] = note
	return note
}

func (ns *NoteStore) GetAll() []models.Note {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	notes := make([]models.Note, 0, len(ns.notes))
	for _, note := range ns.notes {
		notes = append(notes, note)
	}
	return notes
}

func (ns *NoteStore) GetByID(id int) (models.Note, bool) {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	note, exists := ns.notes[id]
	return note, exists
}

func (ns *NoteStore) Update(id int, updatedNote models.Note) (models.Note, bool) {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	if _, exists := ns.notes[id]; !exists {
		return updatedNote, false
	}

	updatedNote.ID = id
	ns.notes[id] = updatedNote
	return updatedNote, true
}

func (ns *NoteStore) Delete(id int) bool {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	if _, exists := ns.notes[id]; !exists {
		return false
	}

	delete(ns.notes, id)
	return true
}
