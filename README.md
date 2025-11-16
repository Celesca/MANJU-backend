# Elysia with Bun runtime

## Getting Started
To get started with this template, simply paste this command into your terminal:
```bash
bun create elysia ./elysia-example
```

## Development
To start the development server run:
```bash
bun run dev
```

Open http://localhost:3000/ with your browser to see the result.

This repository contains a simple Elysia.js server with an in-memory `User` model and full CRUD endpoints.

## Endpoints
- GET / - health check
- POST /users - create a new user
- GET /users - list users
- GET /users/:id - get user by id
- PATCH /users/:id - update user
- DELETE /users/:id - delete user

## Run (with Bun)
```bash
bun install
bun run --watch src/index.ts
```

You can test the endpoints with curl. Example:
```bash
curl --header "Content-Type: application/json" \
	--request POST \
	--data '{"email":"test@example.com","name":"Test"}' \
	http://localhost:3000/users
```

The implementation currently uses an in-memory store. For persistent storage, replace `src/repositories/userRepository.ts` with a DB-backed repository (Prisma, TypeORM, etc.).