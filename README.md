# Go Web Application Template

A modern Go web application template that provides a solid foundation for building scalable web applications with REST API.

## Features

- **Core Functionality**

  - JWT-based authentication and authorization
  - User management system
  - Role-based access control (RBAC)
  - Admin dashboard capabilities
  - Internationalization (i18n) support
  - **Common API**: health, auth, i18n, profile (shared between admin and user)
  - Code generator for layers (repository, schema, datatable, service, handler)

- **API Architecture**

  - RESTful API with Fiber
  - OpenAPI/Swagger documentation
  - Request/Response validation
  - Clean and maintainable code structure
  - Strongly typed API contracts

- **Domain-Driven Design**

  - Clean architecture
  - Separation of concerns
  - Domain models with business logic
  - Repository pattern for data access

- **Database**

  - PostgreSQL with GORM
  - Migrations for schema management
  - Soft deletes
  - Optimized queries with preloading

- **Security**
  - JWT token-based authentication
  - Password hashing with bcrypt
  - CORS configuration
  - Rate limiting
  - Helmet security headers

## Tech Stack

| Component                   | Technology              | Notes                              |
| --------------------------- | ----------------------- | ---------------------------------- |
| Web Framework               | Fiber                   | Fast, lightweight, and easy to use |
| Request/Response Validation | go-playground/validator | Standard Go validation             |
| OpenAPI Docs                | swaggo/swag             | Auto-generated Swagger UI          |
| Config Management           | spf13/viper             | Read from .env, config.yml, etc.   |
| Logging                     | uber-go/zap             | High-performance logging           |
| Database ORM                | gorm                    | Popular and easy to use            |
| DB Migration                | golang-migrate          | Versioned migrations               |
| Task Scheduler              | robfig/cron             | Time-based job scheduling          |

## Project Structure

```
.
├── cmd/                    # Command line entry points (admin, user, migrations)
├── internal/               # Private application code (main backend logic)
│   ├── api/                # API routers (grouped by common, admin, user)
│   ├── config/             # Application configuration logic
│   ├── cron_job/           # Scheduled/recurring jobs (cron tasks)
│   ├── datatable/          # Datatable logic for listing/filtering data
│   ├── handler/            # HTTP handlers (grouped by common, admin, user)
│   ├── middleware/         # HTTP middleware (auth, logging, etc.)
│   ├── model/              # Data models (GORM models, domain entities)
│   ├── repository/         # Data access layer (DB queries, repository pattern)
│   ├── routes/             # Route registration (admin, user)
│   ├── schema/             # Request/response schemas (DTOs, validation)
│   ├── server/             # Server startup and configuration
│   ├── service/            # Business logic/services
│   └── validator/          # Custom request validation logic
├── pkg/                    # Public library code (reusable across projects)
│   ├── auth/               # Authentication helpers/utilities
│   ├── config/             # Shared config utilities
│   ├── core/               # Core utilities and helpers
│   ├── datatable/          # Datatable utilities
│   ├── database/           # Database connection and helpers
│   ├── generator/          # Code generator for CRUD layers
│   ├── i18n/               # Internationalization helpers
│   ├── logger/             # Logging utilities
│   ├── middleware/         # Shared middleware utilities
│   ├── response/           # Standardized API response helpers
│   ├── schema/             # Shared schema utilities
│   └── server/             # Server utilities
├── api/                    # (Empty or reserved for future API definitions)
├── docs/                   # Swagger/OpenAPI documentation
├── migrations/             # Database migration files (SQL)
├── i18n/                   # Localization files (translations, locales)
├── frontend/               # Frontend application (Nuxt.js, UI code)
├── config.yaml             # Main configuration file (YAML)
├── docker-compose.yaml     # Docker Compose services definition
├── Makefile                # Build and development commands
├── .env.example            # Example environment variables
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── main.go                 # Main entry point (if not using cmd/)
└── README.md               # Project documentation
```

## Getting Started

### Prerequisites

Before starting development, you need to install the following tools:

```bash
# Install Swagger for API documentation
go install github.com/swaggo/swag/cmd/swag@latest

# Install golang-migrate for database migrations
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/go-template.git
cd go-template
```

2. Install dependencies:

```bash
make deps
```

3. Set up environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Start infrastructure:

```bash
make start-infra
```

5. Run database migrations:

```bash
make migrate-up
```

6. Generate Swagger documentation:

```bash
make gen-swagger
```

### Configuration

The application can be configured using either environment variables or a `config.yaml` file. Here's an example configuration:

```yaml
app:
  name: "go-template"
  version: "1.0.0"
  mode: "development"

server:
  port: 8080
  read_timeout: 60
  write_timeout: 60
  idle_timeout: 120

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "postgres"
  dbname: "go_template"
  sslmode: "disable"

jwt:
  secret: "your-secret-key"
  expiration_period: 24 # hours

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0

log:
  level: "debug"
  filename: "logs/app.log"
  max_size: 100
  max_backups: 3
  max_age: 7
  compress: true
```

## Running the Application

### Development

```bash
make run
```

### Production

```bash
make build
./bin/app
```

## API Documentation

Once the application is running, you can access the Swagger documentation at:

```
http://localhost:8080/swagger/
```

## Routing & Common API

- **Admin routes**: Register via `routes.NewAdminRouter().Register(app)`
- **User routes**: Register via `routes.NewUserRouter().Register(app)`
- **Common routes (health, auth, i18n, profile)** are automatically available in both admin and user APIs.
- No need to pass config or service into the router, all dependencies are instantiated automatically inside.

## Code Generator

- See details at [`pkg/generator/README.md`](pkg/generator/README.md)
- Supports code generation for repository, schema, datatable, service, and handler from a single Go model.

## Available Make Commands

- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run tests
- `make clean` - Clean build files
- `make gen-swagger` - Generate Swagger documentation
- `make migrate-up` - Run database migrations up
- `make migrate-down` - Run database migrations down
- `make start-infra` - Start infrastructure services
- `make stop-infra` - Stop infrastructure services
- `make deps` - Install dependencies
- `make install-tools` - Install development tools

## API Endpoints

### Authentication

- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/refresh` - Refresh JWT token
- `POST /api/v1/auth/logout` - User logout

### Users

- `GET /api/v1/users` - List users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### WebSocket

- `GET /ws` - WebSocket endpoint for real-time communication

### Documentation

- `GET /swagger/*` - Swagger UI documentation
- `GET /health` - Health check endpoint

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Development Setup

### Prerequisites

Before starting development, you need to install the following tools:

```bash
# Install gRPCurl for testing gRPC endpoints
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# Install protoc-gen-go for generating Go code from protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install protoc-gen-connect-go for generating Connect-Go code from protobuf
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# Install protoc-gen-go-grpc for generating gRPC code from protobuf
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install protoc-gen-validate for generating validation code from protobuf
go install github.com/envoyproxy/protoc-gen-validate@latest

# Install golangci-lint for code linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

These tools are required for:

- `grpcurl`: Testing gRPC endpoints
- `protoc-gen-go`: Generating Go code from Protocol Buffers
- `protoc-gen-connect-go`: Generating Connect-Go code from Protocol Buffers
- `protoc-gen-go-grpc`: Generating gRPC code from Protocol Buffers
- `protoc-gen-validate`: Generating validation code from Protocol Buffers
- `golangci-lint`: Code linting and static analysis
