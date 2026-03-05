# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go-based REST API backend for an Inventory/Transaction Management System using Fiber v2 as the web framework and PostgreSQL as the database.

## Technology Stack

| Component | Technology |
|-----------|------------|
| **Module Name** | `bangkho.dev/tcl/test/backend` |
| **Go Version** | 1.25.0 |
| **Web Framework** | Fiber v2 (github.com/gofiber/fiber/v2 v2.52.12) |
| **ORM** | GORM (gorm.io/gorm v1.31.1) |
| **Database Driver** | PostgreSQL (github.com/lib/pq v1.11.2) |
| **Validation** | go-playground/validator/v10 v10.30.1 |
| **JWT** | golang-jwt/jwt/v5 v5.3.1 |
| **Password Hashing** | golang.org/x/crypto/bcrypt |
| **Env Loading** | joho/godotenv v1.5.1 |

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

## Project Structure

```
backend/
├── main.go                          # Entry point
├── go.mod                           # Go module definition
├── Dockerfile                       # Docker build file
├── docker-compose.yaml              # Docker Compose configuration
├── .env                             # Environment variables
├── CLAUDE.md                        # Claude Code guidance
├── app/bin                          # Compiled binary
├── db/
│   ├── db.go                        # Database singleton connection
│   └── migration/                   # Database migrations
│       ├── migration_user.go
│       ├── migration_customer.go
│       ├── migration_inventory.go
│       └── migration_transaction.go
├── helpers/
│   ├── error.go                     # HTTP response helpers
│   └── validator.go                 # Validation helpers
└── modules/
    ├── user/                        # User authentication & roles
    │   ├── model.go                 # User data model
    │   ├── dto.go                   # Data transfer objects
    │   ├── router.go                # Route definitions
    │   ├── handler.go               # HTTP handlers
    │   ├── repository.go            # Data access layer
    │   └── service.go               # Business logic
    ├── customer/                    # Customer management
    │   ├── model.go
    │   ├── dto.go
    │   ├── router.go
    │   ├── handler.go
    │   ├── repository.go
    │   └── service.go
    ├── inventory/                   # Product/SKU inventory
    │   ├── model.go
    │   ├── dto.go
    │   ├── router.go
    │   ├── handler.go
    │   ├── repository.go
    │   └── service.go
    └── transaction/                 # Transaction handling
        ├── model.go
        ├── dto.go
        ├── router.go
        ├── handler.go
        ├── repository.go
        └── service.go
```

## API Endpoints

All routes are prefixed with `/api`. Health check at `/check`.

### User Module (`/api/users`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users/register` | Register new user |
| POST | `/api/users/login` | User login (returns JWT token) |
| GET | `/api/users` | List users (paginated) |
| GET | `/api/users/:id` | Get user by ID |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Delete user |

**Roles**: `admin`, `superuser` (default)

### Customer Module (`/api/customers`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/customers` | Create customer |
| GET | `/api/customers` | List customers (paginated) |
| GET | `/api/customers/:id` | Get customer by ID |
| PUT | `/api/customers/:id` | Update customer |
| DELETE | `/api/customers/:id` | Delete customer |

### Inventory Module (`/api/inventory`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/inventory` | Create inventory item |
| GET | `/api/inventory` | List inventory (paginated) |
| GET | `/api/inventory/:id` | Get item by ID |
| PUT | `/api/inventory/:id` | Update item |
| DELETE | `/api/inventory/:id` | Delete item |

**Fields**: SKU (unique), Name, Quantity, Price

### Transaction Module (`/api/transactions`)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/transactions` | Create transaction |
| GET | `/api/transactions` | List transactions (paginated) |
| GET | `/api/transactions/:id` | Get transaction by ID |
| PUT | `/api/transactions/:id` | Update transaction |
| DELETE | `/api/transactions/:id` | Delete transaction |

**Transaction Types**: `in`, `out`
**Statuses**: `progress`, `done`, `cancelled`

## Architecture

**Pattern**: Repository pattern with service layer
- Each module has: model, dto, router, handler, repository, service
- DTOs for request/response separation
- Pagination support on all list endpoints (`page`, `page_size` query params)
- Centralized error handling in `helpers/error.go`
- Validation via go-playground/validator with struct tags

**Database Layer**: Singleton connection in `db/db.go` using GORM AutoMigrate

**Entry Point**: `main.go` initializes database, creates Fiber app, and registers routes

## Configuration

Environment variables loaded from `.env` file using godotenv:

```
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=...
POSTGRES_DB=inventory
POSTGRES_PORT=5432
```

Database credentials are also managed through Docker Compose.

## Notes

- JWT token generated on login has 24-hour expiry
- Password hashing using bcrypt
- All list endpoints support pagination via `page` and `page_size` query parameters
- Transaction module validates customer and inventory existence
- Transaction "out" type validates sufficient inventory quantity
