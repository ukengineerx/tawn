# Tawn

A blank platform skeleton: one mobile app, one API, one admin app, and Postgres. No product domain is filled in. The sample resource is a note, so the database, API, and admin have one path to exercise.

This layout is inspired by a larger services monorepo, and it deliberately leaves out that repo's extra services, Kafka, Redis, API gateway, realtime server, and AWS modules. See [docs/infrastructure.md](docs/infrastructure.md).

## Layout

```
├── backend/                 # one Go API (Gin + pgx)
├── deployments/local/       # Postgres only
├── mobile/app/              # Expo (iOS and Android)
├── web/                     # Next.js admin + @tawn/ui
├── docs/infrastructure.md   # what to add when this outgrows one machine
└── .github/workflows/ci.yml
```

The API talks to Postgres through a `notes.Store` interface. A later split into more services can move that package out without rewriting the HTTP handlers.

## Prerequisites

- Go 1.26
- Docker
- Node 22 or 24 for the mobile app (Expo SDK 58 does not support Node 25). The admin app can use the Node already on this machine.

## Run it

```bash
make dev-up      # Postgres on localhost:5433
make run-api     # http://localhost:8080
make web-dev     # http://localhost:3000
make mobile-start
```

Install JS dependencies once:

```bash
npm install --prefix web
npm install --prefix mobile/app
```

| Check | What it tells you |
|-------|-------------------|
| `GET /health` | The process is up. It does not touch the database. |
| `GET /ready` | Postgres answered. |
| `GET /v1/notes` and `POST /v1/notes` | The sample resource. Body is `{"body":"hello"}`. |

`make check` runs the Go tests, Go vet, the admin typecheck and lint, and the Expo typecheck.
