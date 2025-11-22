// mini-http-server/main.go
package main

import (
	"log"
	"mini-http-server/handlers"
	"mini-http-server/utils"
	"net/http"
	"time"
)

func main() {
	// Initialize the note handler
	noteHandler := handlers.NewNoteHandler()

	// Setup routes
	mux := http.NewServeMux()

	// Note routes
	mux.HandleFunc("GET /notes", noteHandler.GetAllNotes)
	mux.HandleFunc("GET /notes/{id}", noteHandler.GetNoteByID)
	mux.HandleFunc("POST /notes", noteHandler.CreateNote)
	mux.HandleFunc("PUT /notes/{id}", noteHandler.UpdateNote)
	mux.HandleFunc("DELETE /notes/{id}", noteHandler.DeleteNote)

	// Health check route
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Home route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{
			"message":   "Mini HTTP Server API",
			"version":   "1.0.0",
			"endpoints": "GET /notes, POST /notes, GET /notes/{id}, PUT /notes/{id}, DELETE /notes/{id}, GET /health",
		})
	})

	// Configure server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	log.Printf("Server starting on http://localhost:8080")
	log.Printf("Available endpoints:")
	log.Printf("  GET  /             - API information")
	log.Printf("  GET  /health       - Health check")
	log.Printf("  GET  /notes        - List all notes")
	log.Printf("  POST /notes        - Create a note")
	log.Printf("  GET  /notes/{id}   - Get a specific note")
	log.Printf("  PUT  /notes/{id}   - Update a note")
	log.Printf("  DELETE /notes/{id} - Delete a note")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
