# CRUD App Tasks

A small Go REST API for managing tasks. The app uses [Fiber v3](https://gofiber.io/) for HTTP routing, [GORM](https://gorm.io/) for persistence, and PostgreSQL for storage.

## Features

- Create a task
- List all tasks
- Get a task by ID
- Update a task by ID
- Delete a task by ID
- Automatically migrates the task table on startup

## Tech Stack

- Go 1.26.2
- Fiber v3
- GORM
- PostgreSQL
- godotenv

## Project Structure

```text
.
├── api/
│   ├── create.go
│   ├── delete.go
│   ├── read.go
│   └── update.go
├── db/
│   └── db.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Prerequisites

- Go installed
- A PostgreSQL database
- A database connection string

## Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://USER:PASSWORD@HOST:PORT/DB_NAME
PORT=3000
```

`PORT` is optional. If it is not set, the server runs on port `3000`.

## Installation

Install dependencies:

```bash
go mod download
```

## Running the App

Start the server:

```bash
go run .
```

By default, the API is available at:

```text
http://localhost:3000
```

## API Endpoints

### Create a Task

```http
POST /tasks/create
```

Example request:

```bash
curl -X POST http://localhost:3000/tasks/create \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Write README",
    "description": "Document the task CRUD API",
    "task": "Create project documentation"
  }'
```

### Get All Tasks

```http
GET /tasks
```

Example request:

```bash
curl http://localhost:3000/tasks
```

### Get One Task

```http
GET /tasks/:id
```

Example request:

```bash
curl http://localhost:3000/tasks/1
```

### Update a Task

```http
PUT /tasks/update/:id
```

Example request:

```bash
curl -X PUT http://localhost:3000/tasks/update/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated title",
    "description": "Updated description",
    "task": "Updated task"
  }'
```

### Delete a Task

```http
DELETE /tasks/delete/:id
```

Example request:

```bash
curl -X DELETE http://localhost:3000/tasks/delete/1
```

## Task Model

Tasks include GORM's default model fields plus the app-specific fields:

```json
{
  "ID": 1,
  "CreatedAt": "2026-05-10T12:00:00Z",
  "UpdatedAt": "2026-05-10T12:00:00Z",
  "DeletedAt": null,
  "title": "Write README",
  "description": "Document the task CRUD API",
  "task": "Create project documentation"
}
```

## Notes

- The app loads environment variables from `.env` during startup.
- The database schema is migrated automatically with `AutoMigrate`.
- If no tasks exist, `GET /tasks` returns a `404` response with `{"error":"No tasks found"}`.
