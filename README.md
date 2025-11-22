# Mini HTTP Server

RESTful note-taking API using only Go standard library.

## Skills
`net/http`, REST APIs, JSON handling, routing, error handling

## Quick Start
```bash
go run main.go
```
Server runs at http://localhost:8080

## Features
- Full CRUD operations for notes  
- Proper HTTP status codes and JSON responses  
- Thread-safe in-memory storage  
- Clean package structure: handlers, models, utils  

## API Endpoints
- `POST   /notes`      – Create a new note  
- `GET    /notes`      – List all notes  
- `GET    /notes/{id}` – Get a single note  
- `PUT    /notes/{id}` – Update a note  
- `DELETE /notes/{id}` – Delete a note  

## Test Commands
```bash
# Create note
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d '{"title":"Test","content":"Hello"}'

# List notes
curl http://localhost:8080/notes

# Get single note
curl http://localhost:8080/notes/1

# Update note
curl -X PUT http://localhost:8080/notes/1 -H "Content-Type: application/json" -d '{"title":"Updated","content":"World"}'

# Delete note
curl -X DELETE http://localhost:8080/notes/1
```

MIT License