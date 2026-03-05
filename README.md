# Inventory Management API

Go-based REST API backend for an Inventory/Transaction Management System.

## Technology Stack

- **Language**: Go 1.25
- **Web Framework**: Fiber v2
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT (24-hour expiry)
- **Validation**: go-playground/validator

## Architecture

This project follows a **layered architecture** with the **Repository Pattern**:

```
modules/
├── user/         # Authentication & user management
├── customer/     # Customer CRUD operations
├── inventory/    # Product/SKU management
└── transaction/ # Transaction handling (in/out types)
```

Each module follows a consistent structure:

| File | Responsibility |
|------|----------------|
| `model.go` | Database schema and struct definitions |
| `dto.go` | Request/Response data transfer objects |
| `repository.go` | Data access layer (database queries) |
| `service.go` | Business logic layer |
| `handler.go` | HTTP request handling |
| `router.go` | Route definitions |

### Key Design Patterns

- **Singleton Database**: Centralized database connection in `db/db.go`
- **Dependency Injection**: Services receive repositories, handlers receive services
- **DTO Separation**: Clean separation between API contracts and database models
- **Pagination**: All list endpoints support `page` and `page_size` query parameters

## Prerequisites

- Go 1.25+
- PostgreSQL 12+ (or Docker)
- Docker & Docker Compose (for containerized deployment)

## Environment Variables

Create a `.env` file in the root directory:

```env
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_password
POSTGRES_DB=inventory
POSTGRES_PORT=5432
```

---

## Development

### Option 1: Local Development

Run PostgreSQL using Docker, then start the Go application locally.

```bash
# 1. Start PostgreSQL container
docker-compose up db

# 2. Run the application
go run main.go
```

The server will start at `http://localhost:8000`

### Option 2: Using Air (Live Reload)

For development with auto-reload:

```bash
# Install air
go install github.com/air-verse/air@latest

# Run with air
air
```

---

## Production

### Using Docker Compose

Build and run all services (backend + database):

```bash
# Build and start all containers
docker-compose up --build

# Run in detached mode
docker-compose up --build -d
```

### Container Services

| Service | Port | Description |
|---------|------|-------------|
| backend | 8000 | REST API server |
| db | 5432 | PostgreSQL database |

### Managing Containers

```bash
# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Stop and remove volumes (database data)
docker-compose down -v
```

### Rebuilding

```bash
# Force rebuild without cache
docker-compose build --no-cache

# Rebuild and start
docker-compose up --build
```

---

## API Endpoints

### Health Check

```
GET /check
```

### User Module

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/users/register` | Register new user |
| POST | `/api/users/login` | Login (returns JWT) |
| GET | `/api/users` | List users (paginated) |
| GET | `/api/users/:id` | Get user by ID |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Delete user |

### Customer Module

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/customers` | Create customer |
| GET | `/api/customers` | List customers (paginated) |
| GET | `/api/customers/:id` | Get customer by ID |
| PUT | `/api/customers/:id` | Update customer |
| DELETE | `/api/customers/:id` | Delete customer |

### Inventory Module

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/inventory` | Create item |
| GET | `/api/inventory` | List items (paginated) |
| GET | `/api/inventory/:id` | Get item by ID |
| PUT | `/api/inventory/:id` | Update item |
| DELETE | `/api/inventory/:id` | Delete item |

### Transaction Module

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/transactions` | Create transaction |
| GET | `/api/transactions` | List transactions (paginated) |
| GET | `/api/transactions/:id` | Get transaction by ID |
| PUT | `/api/transactions/:id` | Update transaction |
| DELETE | `/api/transactions/:id` | Delete transaction |

**Transaction Types**: `in`, `out`
**Statuses**: `progress`, `done`, `cancelled`

---

## Pagination

All list endpoints support pagination:

```
GET /api/inventory?page=1&page_size=10
```

| Parameter | Default | Description |
|-----------|---------|-------------|
| page | 1 | Page number |
| page_size | 10 | Items per page |

---

## Example Usage

### Register a User

```bash
curl -X POST http://localhost:8000/api/users/register \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password123", "role": "admin"}'
```

### Login

```bash
curl -X POST http://localhost:8000/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password123"}'
```

### Create Inventory Item

```bash
curl -X POST http://localhost:8000/api/inventory \
  -H "Content-Type: application/json" \
  -d '{"sku": "SKU001", "name": "Product A", "quantity": 100, "price": 9.99}'
```

### Create Transaction

```bash
curl -X POST http://localhost:8000/api/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "inventory_id": 1,
    "transaction_type": "out",
    "quantity": 5,
    "status": "progress"
  }'
```
