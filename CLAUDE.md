# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based REST API backend for an Inventory/Transaction Management System using Fiber v2 as the web framework and PostgreSQL as the database.

## Technology Stack

- **Web Framework**: Fiber v2
- **ORM**: GORM (gorm.io/gorm)
- **Database**: PostgreSQL
- **Validation**: go-playground/validator (github.com/go-playground/validator/v10)

## Common Commands

```bash
# Start PostgreSQL via Docker
docker-compose up db

# Run the application locally
go run main.go

# Build and run everything with Docker Compose
docker-compose up --build

# Build the Go binary
go build -o app/bin main.go
```

The server runs on port **8000** with a health check endpoint at `GET /check`.

## Architecture

**Modular structure** under `modules/` directory - each domain entity has its own package containing `model.go` with data structures:
- `modules/user/` - User authentication and roles (admin, superuser)
- `modules/customer/` - Customer management
- `modules/inventory/` - Product/SKU inventory tracking
- `modules/transaction/` - Transaction handling (in/out types, progress/done/cancelled statuses)

**Database layer** in `db/` directory - singleton connection pattern in `db/db.go`, migrations in `db/migration/`.

**Entry point**: `main.go` initializes the database, creates the Fiber app, and registers routes.

## Configuration

Environment variables are loaded from `.env` file using godotenv. Database credentials are managed through Docker Compose.
