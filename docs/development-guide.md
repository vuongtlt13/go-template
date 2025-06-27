# Development Guide

## 🚀 Getting Started

This guide will help you set up and develop with the Go web application template.

## 📋 Prerequisites

### Required Software

- **Go 1.23+**: [Download Go](https://golang.org/dl/)
- **PostgreSQL 16+**: [Download PostgreSQL](https://www.postgresql.org/download/)
- **Redis 7+**: [Download Redis](https://redis.io/download)
- **Docker & Docker Compose**: [Download Docker](https://www.docker.com/products/docker-desktop)
- **Node.js 18+**: [Download Node.js](https://nodejs.org/)

### Development Tools

```bash
# Install Swagger for API documentation
go install github.com/swaggo/swag/cmd/swag@latest

# Install golang-migrate for database migrations
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install additional Go tools
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

## 🏗️ Project Setup

### 1. Clone Repository

```bash
git clone <repository-url>
cd go-template
```

### 2. Install Dependencies

```bash
# Backend dependencies
make deps

# Frontend dependencies
cd frontend && yarn install
```

### 3. Environment Configuration

```bash
# Copy environment template
cp .env.example .env

# Edit environment variables
nano .env
```

**Key Environment Variables:**

```bash
# Application
APP_MODE=development
SQL_DEBUG=true

# Server
SERVER_PORT=8000
SERVER_CORS=*

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_template

# JWT
JWT_SECRET=your-secret-key-here
JWT_EXPIRE_PERIOD=24h

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

### 4. Start Infrastructure

```bash
# Start PostgreSQL and Redis
make start-infra

# Verify services are running
docker ps
```

### 5. Database Setup

```bash
# Run database migrations
make migrate-up

# Verify tables created
psql -h localhost -U postgres -d go_template -c "\dt"
```

### 6. Generate Documentation

```bash
# Generate Swagger documentation
make gen-swagger
```

## 🏃‍♂️ Running the Application

### Development Mode

#### Backend Servers

```bash
# Terminal 1: Admin Server
go run main.go admin

# Terminal 2: User Server (optional)
go run main.go user
```

#### Frontend Development

```bash
# Terminal 3: Frontend
cd frontend && yarn dev
```

### Production Mode

```bash
# Build the application
make build

# Run with production config
APP_MODE=production ./bin/app admin
```

## 🛠️ Development Workflow

### 1. Code Structure

#### Adding New Features

```
internal/
├── model/          # Define your entity
├── repository/     # Data access layer
├── service/        # Business logic
├── handler/        # HTTP handlers
├── schema/         # Request/Response DTOs
└── datatable/      # Data listing logic
```

#### Example: Adding a Product Feature

**Step 1: Define Model**

```go
// internal/model/product.go
type Product struct {
    ID          uint64         `gorm:"primarykey" json:"id"`
    Name        string         `gorm:"size:255;not null" json:"name"`
    Description string         `gorm:"type:text" json:"description"`
    Price       decimal.Decimal `gorm:"type:decimal(10,2)" json:"price"`
    IsActive    bool           `gorm:"default:true" json:"is_active"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Step 2: Generate Code**

```bash
# Generate all layers for Product
go run pkg/generator/main.go --cmd=generate --model=Product --component=repository
go run pkg/generator/main.go --cmd=generate --model=Product --component=service
go run pkg/generator/main.go --cmd=generate --model=Product --component=handler
go run pkg/generator/main.go --cmd=generate --model=Product --component=schema
go run pkg/generator/main.go --cmd=generate --model=Product --component=datatable
```

**Step 3: Create Migration**

```bash
# Create migration file
migrate create -ext sql -dir migrations -seq create_products_table
```

**Step 4: Implement Migration**

```sql
-- migrations/000002_create_products_table.up.sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_products_name ON products(name);
CREATE INDEX idx_products_is_active ON products(is_active);

-- migrations/000002_create_products_table.down.sql
DROP TABLE products;
```

**Step 5: Register Routes**

```go
// internal/api/admin/product.go
package admin

import (
    "github.com/gofiber/fiber/v2"
    "yourapp/internal/handler/admin"
)

type ProductRouter struct{}

func NewProductRouter() *ProductRouter {
    return &ProductRouter{}
}

func (r *ProductRouter) Register(router fiber.Router) {
    productHandler := admin.NewProductHandler()

    products := router.Group("/products")
    products.Get("/", productHandler.List)
    products.Get("/:id", productHandler.Get)
    products.Post("/", productHandler.Create)
    products.Put("/:id", productHandler.Update)
    products.Delete("/:id", productHandler.Delete)
    products.Post("/datatable", productHandler.DataTable)
}
```

**Step 6: Update Main Router**

```go
// internal/routes/admin.go
func (r *AdminRouter) Register(app *fiber.App) {
    // ... existing code ...

    // Product management
    productRouter := admin.NewProductRouter()
    productRouter.Register(api)
}
```

### 2. Testing

#### Unit Tests

```go
// internal/service/product_test.go
func TestProductService_Create(t *testing.T) {
    // Setup
    db := setupTestDB()
    repo := repository.NewProductRepository(db)
    service := NewProductService(db, repo)

    // Test
    product := &model.Product{
        Name:        "Test Product",
        Description: "Test Description",
        Price:       decimal.NewFromFloat(99.99),
        IsActive:    true,
    }

    err := service.Create(context.Background(), product)

    // Assert
    assert.NoError(t, err)
    assert.NotZero(t, product.ID)
}
```

#### Integration Tests

```go
// internal/handler/admin/product_test.go
func TestProductHandler_Create(t *testing.T) {
    // Setup
    app := fiber.New()
    handler := NewProductHandler()

    app.Post("/products", handler.Create)

    // Test
    req := httptest.NewRequest("POST", "/products", strings.NewReader(`{
        "name": "Test Product",
        "description": "Test Description",
        "price": 99.99
    }`))
    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
```

#### Running Tests

```bash
# Run all tests
make test

# Run specific test
go test ./internal/service -v

# Run with coverage
go test ./... -cover
```

### 3. Code Quality

#### Linting

```bash
# Run linter
golangci-lint run

# Fix issues automatically
golangci-lint run --fix
```

#### Formatting

```bash
# Format code
go fmt ./...

# Organize imports
goimports -w .
```

#### Pre-commit Hooks

```bash
# Install pre-commit hooks
cp .git/hooks/pre-commit.sample .git/hooks/pre-commit

# Edit pre-commit hook
nano .git/hooks/pre-commit
```

**Pre-commit Hook Content:**

```bash
#!/bin/sh
# Pre-commit hook

echo "Running tests..."
go test ./...

echo "Running linter..."
golangci-lint run

echo "Formatting code..."
go fmt ./...
goimports -w .

echo "Pre-commit checks passed!"
```

## 🔧 Configuration Management

### Environment-Specific Configs

#### Development

```yaml
# config.dev.yaml
app:
  mode: "development"
  debug: true

server:
  port: 8000
  cors: "*"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "postgres"
  dbname: "go_template_dev"
```

#### Production

```yaml
# config.prod.yaml
app:
  mode: "production"
  debug: false

server:
  port: 8080
  cors: "https://yourdomain.com"

database:
  host: "prod-db-host"
  port: 5432
  user: "prod_user"
  password: "${DB_PASSWORD}"
  dbname: "go_template_prod"
```

### Configuration Loading

```go
// Load config based on environment
func loadConfig() (*config.Config, error) {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }

    viper.SetConfigName(fmt.Sprintf("config.%s", env))
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    return config.GetConfig(), nil
}
```

## 🗄️ Database Development

### Migration Workflow

#### Creating Migrations

```bash
# Create new migration
migrate create -ext sql -dir migrations -seq add_user_profile_fields

# This creates:
# - 000003_add_user_profile_fields.up.sql
# - 000003_add_user_profile_fields.down.sql
```

#### Migration Best Practices

```sql
-- Up migration
-- 000003_add_user_profile_fields.up.sql

-- Add new columns
ALTER TABLE users
ADD COLUMN phone VARCHAR(20),
ADD COLUMN address TEXT,
ADD COLUMN date_of_birth DATE;

-- Add indexes for performance
CREATE INDEX idx_users_phone ON users(phone);

-- Add constraints
ALTER TABLE users
ADD CONSTRAINT chk_phone_format
CHECK (phone ~ '^\+?[1-9]\d{1,14}$');

-- Down migration
-- 000003_add_user_profile_fields.down.sql

-- Remove constraints
ALTER TABLE users DROP CONSTRAINT IF EXISTS chk_phone_format;

-- Remove indexes
DROP INDEX IF EXISTS idx_users_phone;

-- Remove columns
ALTER TABLE users
DROP COLUMN IF EXISTS phone,
DROP COLUMN IF EXISTS address,
DROP COLUMN IF EXISTS date_of_birth;
```

### Database Seeding

```go
// internal/database/seed.go
func SeedDatabase(db *gorm.DB) error {
    // Seed roles
    roles := []model.Role{
        {Code: "admin", Name: "Administrator", Description: "Full system access"},
        {Code: "user", Name: "User", Description: "Regular user access"},
    }

    for _, role := range roles {
        if err := db.FirstOrCreate(&role, model.Role{Code: role.Code}).Error; err != nil {
            return err
        }
    }

    // Seed permissions
    permissions := []model.Permission{
        {Code: "user:create", Name: "Create Users", Service: "user", Method: "create"},
        {Code: "user:read", Name: "View Users", Service: "user", Method: "read"},
        // ... more permissions
    }

    for _, perm := range permissions {
        if err := db.FirstOrCreate(&perm, model.Permission{Code: perm.Code}).Error; err != nil {
            return err
        }
    }

    return nil
}
```

## 🔍 Debugging

### Logging

```go
// Structured logging
logger := zap.NewProduction()
defer logger.Sync()

logger.Info("User created",
    zap.String("email", user.Email),
    zap.Uint64("user_id", user.ID),
    zap.String("ip_address", c.IP()),
)
```

### Debug Mode

```bash
# Enable debug mode
SQL_DEBUG=true go run main.go admin

# Enable verbose logging
LOG_LEVEL=debug go run main.go admin
```

### Database Debugging

```sql
-- Enable query logging
SET log_statement = 'all';
SET log_min_duration_statement = 0;

-- View slow queries
SELECT query, calls, total_time, mean_time
FROM pg_stat_statements
ORDER BY mean_time DESC
LIMIT 10;
```

## 🚀 Deployment

### Docker Development

```dockerfile
# Dockerfile.dev
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main", "admin"]
```

```yaml
# docker-compose.dev.yml
version: "3.8"
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile.dev
    ports:
      - "8000:8000"
    environment:
      - APP_MODE=development
      - DB_HOST=postgres
      - REDIS_HOST=redis
    depends_on:
      - postgres
      - redis
    volumes:
      - .:/app
      - /app/vendor
```

### Production Deployment

```bash
# Build production image
docker build -t go-template:latest .

# Run with production config
docker run -d \
  --name go-template \
  -p 8080:8080 \
  -e APP_MODE=production \
  -e DB_HOST=prod-db \
  go-template:latest
```

## 📚 Best Practices

### Code Organization

- Follow Go naming conventions
- Use meaningful variable and function names
- Keep functions small and focused
- Use interfaces for dependency injection

### Error Handling

```go
// Always handle errors
if err != nil {
    logger.Error("Failed to create user", zap.Error(err))
    return response.ErrorResponse(c, fiber.StatusInternalServerError, "Internal server error", 500)
}

// Use custom error types
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidEmail = errors.New("invalid email format")
)
```

### Security

- Always validate input
- Use parameterized queries
- Implement proper authentication
- Follow OWASP guidelines

### Performance

- Use connection pooling
- Implement caching strategies
- Optimize database queries
- Monitor application metrics

## 🆘 Troubleshooting

### Common Issues

#### Database Connection

```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Test connection
psql -h localhost -U postgres -d go_template -c "SELECT 1;"

# Check logs
docker logs go_template_postgres
```

#### Migration Issues

```bash
# Check migration status
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/go_template?sslmode=disable" version

# Force migration to specific version
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/go_template?sslmode=disable" force 1
```

#### Port Conflicts

```bash
# Check what's using port 8000
lsof -i :8000

# Kill process
kill -9 <PID>
```

### Getting Help

- Check the logs: `docker logs <container_name>`
- Review the documentation
- Check GitHub issues
- Ask in the community

This development guide provides comprehensive instructions for setting up, developing, and deploying the Go web application template.
