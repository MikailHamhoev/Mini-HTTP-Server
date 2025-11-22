// mini-http-server/handlers/handlers.go
package handlers

import (
	"encoding/json"
	"mini-http-server/models"
	"mini-http-server/utils"
	"net/http"
	"strconv"
	"time"
)

type NoteHandler struct {
	store *NoteStore
}

func NewNoteHandler() *NoteHandler {
	return &NoteHandler{
		store: NewNoteStore(),
	}
}

func (h *NoteHandler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := h.store.GetAll()
	utils.RespondWithJSON(w, http.StatusOK, notes)
}

func (h *NoteHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	note, exists := h.store.GetByID(id)
	if !exists {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, note)
}

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req models.CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Title is required")
		return
	}

	now := time.Now()
	note := models.Note{
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdNote := h.store.Create(note)
	utils.RespondWithJSON(w, http.StatusCreated, createdNote)
}

func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	var req models.UpdateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Title is required")
		return
	}

	// Check if note exists
	existingNote, exists := h.store.GetByID(id)
	if !exists {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	updatedNote := models.Note{
		ID:        id,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: existingNote.CreatedAt,
		UpdatedAt: time.Now(),
	}

	if _, success := h.store.Update(id, updatedNote); !success {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update note")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedNote)
}

func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	if !h.store.Delete(id) {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
