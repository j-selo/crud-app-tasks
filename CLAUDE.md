# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Install dependencies
go mod download

# Run the app
go run .

# Build the binary
go build -o crud-app-tasks .
```

No test suite exists yet. The app requires a valid `.env` with `DATABASE_URL` to start — it calls `log.Fatal` on connection failure.

## Architecture

The app is a thin REST API with three layers:

- **`main.go`** — wires up the Fiber app, calls `db.ConnectDB().AutoMigrate(&api.Task{})` on startup, and registers all routes.
- **`db/db.go`** — singleton GORM connection. `ConnectDB()` initializes `db.DB` on first call and returns it on every subsequent call. All handlers call `db.ConnectDB()` directly; there is no injected dependency.
- **`api/`** — one file per CRUD operation. The `Task` struct (with `gorm.Model` embedded) is defined in `create.go` and reused across all other handler files in the same package.

## Routes

| Method | Path | Handler |
|--------|------|---------|
| POST | `/create` | `api.CreateTask` |
| GET | `/` | `api.GetTasks` |
| GET | `/:id` | `api.GetTask` |
| PUT | `/update/:id` | `api.UpdateTask` |
| DELETE | `/delete/:id` | `api.DeleteTask` |

## Environment

```env
DATABASE_URL=postgres://USER:PASSWORD@HOST:PORT/DB_NAME
PORT=3000  # optional, defaults to 3000
```
