# Go Web Application Template - Documentation

A modern, enterprise-grade Go web application template that provides a solid foundation for building scalable web applications with REST API, following Clean Architecture principles and Domain-Driven Design patterns.

## 🏗️ Project Overview

This template implements a **full-stack web application** with a robust backend built in Go and a modern frontend using Nuxt.js 3. The architecture is designed for scalability, maintainability, and enterprise-level requirements.

### Key Features

- **Multi-server Architecture**: Separate admin and user servers
- **Clean Architecture**: Clear separation of concerns across layers
- **JWT Authentication & Authorization**: Secure user management with RBAC
- **Internationalization**: Multi-language support
- **Code Generation**: Automated CRUD code generation
- **API Documentation**: Auto-generated Swagger/OpenAPI docs
- **Database Migration**: Versioned schema management
- **Background Jobs**: Scheduled task processing
- **Modern Frontend**: Nuxt.js 3 with Vuetify and TypeScript

## 🚀 Quick Start

### Prerequisites

```bash
# Install Swagger for API documentation
go install github.com/swaggo/swag/cmd/swag@latest

# Install golang-migrate for database migrations
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Installation

1. **Clone and setup**:

   ```bash
   git clone <repository-url>
   cd go-template
   make deps
   ```

2. **Configure environment**:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start infrastructure**:

   ```bash
   make start-infra  # Starts PostgreSQL and Redis
   ```

4. **Run migrations**:

   ```bash
   make migrate-up
   ```

5. **Generate documentation**:

   ```bash
   make gen-swagger
   ```

6. **Start servers**:

   ```bash
   # Backend
   go run main.go admin    # Admin server on port 8000
   go run main.go user     # User server on port 8001

   # Frontend
   cd frontend && yarn dev
   ```

## 📁 Project Structure

```
go-template/
├── cmd/                    # CLI entry points
│   ├── admin.go           # Admin server command
│   ├── user.go            # User server command
│   └── migrate.go         # Migration command
├── internal/              # Private application code
│   ├── api/               # API routers (admin, user, common)
│   ├── handler/           # HTTP handlers
│   ├── service/           # Business logic layer
│   ├── repository/        # Data access layer
│   ├── model/             # Domain models/entities
│   ├── schema/            # Request/Response DTOs
│   ├── datatable/         # Data listing logic
│   ├── middleware/        # HTTP middleware
│   ├── server/            # Server configuration
│   └── cron_job/          # Scheduled tasks
├── pkg/                   # Public shared libraries
│   ├── auth/              # JWT authentication
│   ├── config/            # Configuration management
│   ├── database/          # Database connection & migration
│   ├── datatable/         # DataTable utilities
│   ├── generator/         # Code generator for CRUD
│   ├── i18n/              # Internationalization
│   ├── logger/            # Logging utilities
│   ├── middleware/        # Shared middleware
│   ├── response/          # Standardized API responses
│   └── server/            # Server utilities
├── frontend/              # Nuxt.js frontend application
├── migrations/            # Database migration files
├── i18n/                  # Translation files
├── docs/                  # Documentation
├── config.yaml            # Main configuration
├── docker-compose.yaml    # Infrastructure services
└── Makefile               # Build and development commands
```

## 🏛️ Architecture Overview

### Multi-Server Architecture

The application supports multiple server instances:

- **Admin Server** (`cmd/admin.go`): Administrative interface and APIs
- **User Server** (`cmd/user.go`): User-facing APIs and services
- **Common APIs**: Shared functionality between servers

### Clean Architecture Layers

```
┌─────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                       │
├─────────────────────────────────────────────────────────────┤
│  Handler (HTTP)  │  Schema (DTO)  │  Validator  │  Router   │
└─────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────┐
│                     BUSINESS LAYER                          │
├─────────────────────────────────────────────────────────────┤
│  Service (Logic)  │  Auth Service  │  User Service  │ ...   │
└─────────────────────────────────────────────────────────────┘
                                │
┌─────────────────────────────────────────────────────────────┐
│                      DATA LAYER                             │
├─────────────────────────────────────────────────────────────┤
│ Repository (DB)  │  Model (Entity)  │  Migration  │  Cache  │
└─────────────────────────────────────────────────────────────┘
```

### Technology Stack

| Component          | Technology              | Purpose                    |
| ------------------ | ----------------------- | -------------------------- |
| **Web Framework**  | Fiber                   | Fast HTTP framework        |
| **Database ORM**   | GORM                    | Object-relational mapping  |
| **Database**       | PostgreSQL              | Primary data store         |
| **Cache**          | Redis                   | Session and data caching   |
| **Authentication** | JWT                     | Token-based authentication |
| **Validation**     | go-playground/validator | Request validation         |
| **Documentation**  | Swagger/OpenAPI         | API documentation          |
| **Configuration**  | Viper                   | Config management          |
| **Logging**        | Uber Zap                | High-performance logging   |
| **Migration**      | golang-migrate          | Database schema management |
| **Task Scheduler** | robfig/cron             | Background job scheduling  |

## 🔧 Core Components

### Configuration Management

The application uses a centralized configuration system with support for environment variables and YAML files:

```go
type Config struct {
    AppMode  string
    Server   ServerConfig
    Database DBConfig
    JWT      JWTConfig
    Redis    RedisConfig
    I18n     I18nConfig
}
```

### Database Layer

PostgreSQL with GORM ORM, featuring:

- Connection pooling
- Soft deletes
- Auto-migration
- Transaction support

### Authentication & Authorization

- **JWT-based authentication** with configurable expiration
- **bcrypt password hashing** for security
- **Role-based access control (RBAC)** with permissions
- **Middleware-based protection** for routes

### Repository Pattern

Generic repository implementation with Go generics:

```go
type BaseRepository[T any] interface {
    Create(ctx context.Context, entity *T, db *gorm.DB) error
    FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*T, error)
    Update(ctx context.Context, entity *T, db *gorm.DB) error
    Delete(ctx context.Context, entity *T, db *gorm.DB) error
    FindAll(ctx context.Context, offset, limit int, preloads []string, db *gorm.DB) ([]T, error)
}
```

### Service Layer

Business logic layer with dependency injection:

```go
type AuthService interface {
    Login(ctx context.Context, cred Credential) (string, error)
    Register(ctx context.Context, cred Credential) error
}
```

### DataTable System

Advanced data listing with:

- Pagination
- Search and filtering
- Export functionality (Excel, CSV, PDF)
- Column customization
- Smart search capabilities

## 🔒 Security Features

### Authentication

- JWT token-based authentication
- Secure password hashing with bcrypt
- Token expiration and refresh mechanisms

### Authorization

- Role-based access control (RBAC)
- Permission-based authorization
- Middleware-based route protection

### Security Headers

- CORS configuration
- Rate limiting
- Helmet security headers
- Input validation and sanitization

## 🌐 Internationalization

Multi-language support with:

- JSON-based translation files
- Context-aware locale detection
- Fallback to default locale
- Dynamic translation loading

## ⚡ Performance Features

### Database Optimization

- Connection pooling
- Query optimization
- Preloading relationships
- Index management

### Caching Strategy

- Redis integration for session storage
- In-memory caching for translations
- Query result caching

### Background Processing

- Scheduled tasks with cron
- Asynchronous job processing
- Resource cleanup jobs

## 🛠️ Development Tools

### Code Generator

Automated code generation for CRUD operations:

```bash
# Generate repository layer
go run pkg/generator/main.go --cmd=generate --model=User --component=repository

# Generate service layer
go run pkg/generator/main.go --cmd=generate --model=User --component=service

# Generate handler layer
go run pkg/generator/main.go --cmd=generate --model=User --component=handler
```

### Database Migration

Versioned database schema management:

```bash
# Run migrations up
make migrate-up

# Run migrations down
make migrate-down
```

### API Documentation

Auto-generated Swagger documentation:

```bash
# Generate documentation
make gen-swagger

# Access at: http://localhost:8000/swagger/
```

## 📊 API Response Format

### Success Response

```json
{
  "success": true,
  "data": { ... },
  "message": "Operation completed successfully"
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error description",
  "code": 400
}
```

## 🚀 Deployment

### Docker Support

The project includes Docker Compose configuration for:

- PostgreSQL database
- Redis cache
- Application containers

### Environment Configuration

- Development, staging, and production configurations
- Environment variable support
- Secure secret management

## 📚 Additional Documentation

- [API Documentation](./api.md) - Detailed API reference
- [Database Schema](./database.md) - Database design and relationships
- [Code Generator](./generator.md) - Code generation guide
- [Deployment Guide](./deployment.md) - Production deployment instructions
- [Development Guide](./development.md) - Development workflow and best practices

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ using Go and modern web technologies**
