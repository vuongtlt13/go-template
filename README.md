# Go Web Application Template

A modern Go web application template that provides a solid foundation for building scalable web applications with gRPC/Connect.

## Features

- **Core Functionality**
  - JWT-based authentication and authorization
  - User management system
  - Role-based access control (RBAC)
  - Admin dashboard capabilities
  - Internationalization (i18n) support

- **API Architecture**
  - Connect gRPC for high-performance RPC
  - Protocol Buffers for API definitions
  - Strongly typed API contracts
  - Code generation from proto files

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

## Project Structure

```
.
├── cmd/                    # Command line applications
│   ├── admin/             # Admin server
│   └── user/              # User server
├── internal/              # Private application code
│   ├── handler/           # HTTP handlers
│   │   ├── admin/        # Admin handlers
│   │   └── user/         # User handlers
│   ├── repository/        # Data access layer
│   ├── service/          # Business logic layer
│   └── server/           # Server implementations
├── pkg/                   # Public library code
│   ├── config/           # Configuration
│   ├── database/         # Database connection
│   ├── i18n/             # Internationalization
│   ├── logger/           # Logging
│   ├── middleware/       # HTTP middleware
│   └── server/           # Base server
└── pb/                    # Protocol buffers
    ├── admin/            # Admin service
    └── auth/             # Auth service
```

## Getting Started

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

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-template.git
cd go-template
```

2. Install dependencies:
```bash
go mod download
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

6. Generate Protocol Buffer code:
```bash
make gen-proto
```

### Environment Variables

Create a `.env` file in the project root:

```env
# App Configuration
APP_MODE=development
SQL_DEBUG=true
CORS=*
DEFAULT_LOCALE=en  # Default language for i18n (e.g., en, vi)

# Server Configuration
SERVER_USER_PORT=8080
SERVER_ADMIN_PORT=8081

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres

# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRE_PERIOD=24h

# Email Configuration
EMAIL_HOST=localhost
EMAIL_PORT=1025
EMAIL_USERNAME=
EMAIL_PASSWORD=
EMAIL_FROM=

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# Job Configuration
JOB_CONCURRENCY=5

# I18n Configuration
I18N_DEFAULT_LOCALE=en
I18N_BASE_FOLDER=i18n/locales
```

## Running the Application

### Admin Server
```bash
go run cmd/admin/main.go
```

### User Server
```bash
go run cmd/user/main.go
```

## Development

### Available Make Commands

- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run tests
- `make clean` - Clean build files
- `make gen-proto` - Generate Protocol Buffer code
- `make migrate-up` - Run database migrations up
- `make migrate-down` - Run database migrations down
- `make start-infra` - Start infrastructure services
- `make stop-infra` - Stop infrastructure services

### Protocol Buffer Generation

The template uses Protocol Buffers for API definitions. To regenerate the code after modifying proto files:

```bash
make gen-proto
```

### Database Migrations

Create new migration files in `migrations/` directory:
```bash
# Apply migrations
make migrate-up

# Rollback migrations
make migrate-down
```

### Using Middleware

The template includes several useful middleware components:

- **i18n Middleware**: Automatically detects the user's preferred language from the `Accept-Language` header and sets it in the request context. If the language is not supported, it falls back to the default locale.

- **Request Middleware**: Logs incoming requests and their details.

### Internationalization (i18n)

The template supports internationalization (i18n) for multiple languages. Translations are stored in JSON files under the `i18n/locales/` directory, organized by language code (e.g., `en`, `vi`).

#### Translation Files Structure

```
i18n/locales/
├── en/
│   ├── auth.json
│   ├── crud.json
│   └── models/
│       └── user.json
└── vi/
    ├── auth.json
    ├── crud.json
    └── models/
        └── user.json
```

#### Using Translations in Code

To use translations in your Go code, import the `i18n` package and call the `T` function:

```go
import "yourapp/pkg/i18n"

func someFunction(ctx context.Context) {
    translatedText := i18n.T(ctx, "auth.login.title")
    // translatedText will be "Login" for English or "Đăng nhập" for Vietnamese
}
```

#### i18n API Endpoint

The i18n service provides an API endpoint to retrieve translations for a specific language:

- **Endpoint**: `GET /lang/{language}.json`
- **Example**: `GET /lang/en.json` returns all English translations.

The response is a JSON-encoded map of translation keys to their values.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

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
