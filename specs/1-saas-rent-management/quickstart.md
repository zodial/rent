# Quickstart: saas-rent-management (local dev)

**Created**: 2025-11-14

This quickstart shows how to run the backend (Go + Gin + GORM + SQLite) and the frontend (Vue 3 + Vite) locally for development.

Prerequisites
- Go 1.20+
- Node 18+ and npm or pnpm
- git

Backend (server)

1. Initialize Go module (if not already):

```bash
cd backend
go mod init github.com/yourorg/rent-backend
# (if go.mod exists, skip)
```

2. Install dependencies and run migrations (GORM AutoMigrate is used for MVP):

```bash
cd backend
# fetch deps
go get -u github.com/gin-gonic/gin github.com/stretchr/testify gorm.io/gorm gorm.io/driver/sqlite github.com/mattn/go-sqlite3
# run migrations (example binary)
go run ./cmd/server --migrate
```

3. Run the server:

```bash
# set environment variables (example)
export RUNTIME_ENV=development
export DATABASE_URL=./data/dev.sqlite
# start server
go run ./cmd/server
```

API will be reachable at http://localhost:8080/api/v1 by default.

Frontend (admin + client)

1. Scaffold frontend (if not present):

```bash
cd frontend
# Use your preferred package manager (npm/pnpm/yarn)
npm install
npm run dev
```

2. Development notes:
- Frontend should use fetch()/native Web APIs for network calls preferring small helper wrappers in `src/services`.
- Ensure vitest is configured for TDD and tests run in watch mode during development.

Testing & CI
- Run backend unit tests:

```bash
cd backend
go test ./... -coverprofile=coverage.out
```

- Run frontend tests:

```bash
cd frontend
npm run test:unit
```

- CI gates should run lint -> unit tests -> integration tests -> coverage check -> performance checks (if available).


Troubleshooting
- If SQLite foreign keys enforcement is needed, ensure PRAGMA foreign_keys=ON is set when opening DB connections.
- For test isolation, use temporary sqlite file paths or :memory: connections where supported.



<!-- End quickstart -->
