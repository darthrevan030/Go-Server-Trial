# go-server-trial

A RESTful API built in Go as a learning project to understand backend fundamentals — clean architecture, MongoDB integration, and HTTP routing.

## Tech Stack

- **Language:** Go 1.25
- **Router:** [chi](https://github.com/go-chi/chi)
- **Database:** MongoDB (via mongo-driver v2)
- **Config:** godotenv

## Project Structure

``` markdown
cmd/go-server-trial/
    main.go                 # Entry point, wires everything together

internal/
    config/
        config.go           # Loads environment variables
    database/
        database.go         # MongoDB connection
    user/
        model.go            # User struct, request/response types
        interface.go        # Repository interface (contract)
        repository.go       # MongoDB implementation
        handler.go          # HTTP handlers
```

## Getting Started

### Prerequisites

- Go 1.25+
- MongoDB running locally on port 27017

### Installation

1. Clone the repository

```bash
git clone https://github.com/darthrevan030/go-server-trial.git
cd go-server-trial
```

1. Install dependencies

```bash
go mod tidy
```

1. Create a `.env.local` file in the project root

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_NAME=go-server-trial
MONGODB_COLLECTION_NAME=users
PORT=3000
```

1. Run the server

```bash
go run ./cmd/go-server-trial/main.go
```

The server will start on `http://localhost:3000`.

## API Endpoints

All routes are prefixed with `/api/v1`.

| Method | Endpoint | Description |
| -------- | ---------- | ------------- |
| GET | `/api/v1/health` | Health check |
| POST | `/api/v1/users` | Create a user |
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/{id}` | Get a user by ID |
| PUT | `/api/v1/users/{id}` | Update a user's age |
| DELETE | `/api/v1/users/{id}` | Delete a user by ID |
| DELETE | `/api/v1/users` | Delete all users |

## Example Requests

### Create a User

```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Veanna", "age": 20, "country": "Singapore"}'
```

### Get All Users

```bash
curl http://localhost:3000/api/v1/users
```

### Get User by ID

```bash
curl http://localhost:3000/api/v1/users/{id}
```

### Update User Age

```bash
curl -X PUT http://localhost:3000/api/v1/users/{id} \
  -H "Content-Type: application/json" \
  -d '{"age": 21}'
```

### Delete User by ID

```bash
curl -X DELETE http://localhost:3000/api/v1/users/{id}
```

## Key Concepts Practiced

- **Feature-based project structure** — all user-related code lives in `internal/user/`
- **Repository pattern** — database logic is isolated behind an interface
- **Dependency injection** — handlers receive the repository, not the DB directly
- **Clean architecture** — handlers don't know about MongoDB, repository doesn't know about HTTP
