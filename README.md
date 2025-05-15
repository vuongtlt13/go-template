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

#### Request Middleware

The request middleware stores the request in the context for easy access throughout the request lifecycle:

```go
// In your handler
func (h *Handler) SomeHandler(c *fiber.Ctx) error {
    // Get request from context
    req := c.Locals("request").(*fiber.Request)
    
    // Access request properties
    headers := req.Header
    body := req.Body()
    params := req.Params
    query := req.QueryArgs()
    cookies := req.Cookies()
    
    // Do something with request data
    ...
}
```

#### I18n Middleware

The i18n middleware handles internationalization based on the `Accept-Language` header:

```go
// In your handler
func (h *Handler) SomeHandler(c *fiber.Ctx) error {
    // Get language from context
    lang := c.Locals("lang").(string)
    
    // Get translation using i18n
    message := i18n.T(c.Context(), "common.hello")
    
    return c.JSON(fiber.Map{
        "message": message,
        "lang": lang,
    })
}
```

Test with curl:
```bash
# English
curl -H "Accept-Language: en" http://localhost:8080/api/hello

# Vietnamese
curl -H "Accept-Language: vi" http://localhost:8080/api/hello

# Unsupported language (falls back to English)
curl -H "Accept-Language: fr" http://localhost:8080/api/hello
```

## Internationalization (i18n)

The template includes a robust internationalization system that supports multiple languages and nested translation structures.

### Directory Structure

Translation files are organized in a hierarchical structure under `i18n/locales/`:

```
i18n/locales/
├── en/                    # English translations
│   ├── auth.json         # Authentication-related translations
│   ├── crud.json         # CRUD operation translations
│   └── models/           # Model-specific translations
│       └── user.json     # User model translations
└── vi/                    # Vietnamese translations
    ├── auth.json
    ├── crud.json
    └── models/
        └── user.json
```

### Translation File Format

Translation files use JSON format with nested structures:

```json
// i18n/locales/en/auth.json
{
  "login": {
    "title": "Login",
    "submit": "Sign In",
    "success": "Login successful",
    "error": "Invalid credentials"
  }
}

// i18n/locales/en/models/user.json
{
  "menu_title": "User",
  "fields": {
    "email": "Email",
    "password": "Password"
  }
}
```

### Configuration

Configure i18n in your `.env` file:

```env
# i18n Configuration
DEFAULT_LOCALE=en  # Default language (e.g., en, vi)
```

### Usage

#### 1. Initialize i18n

```go
import "yourapp/pkg/i18n"
import "yourapp/pkg/config"

// Initialize i18n with configuration
i18n.Init(&config.I18nConfig{
    DefaultLocale: "en",
    BaseFolder:    "i18n/locales",
})
```

#### 2. Using Translations in Handlers

```go
import "yourapp/pkg/i18n"

func (h *Handler) SomeHandler(c *fiber.Ctx) error {
    // Get translation using context
    message := i18n.T(c.Context(), "auth.login.title")
    
    // Get translation with specific locale
    message := i18n.T(c.Context(), "models.user.fields.email")
    
    return c.JSON(fiber.Map{
        "message": message,
    })
}
```

#### 3. Adding New Translations

```go
// Add a new translation programmatically
i18n.AddTranslation("en", "welcome.message", "Welcome to our application")
i18n.AddTranslation("vi", "welcome.message", "Chào mừng đến với ứng dụng của chúng tôi")
```

### Middleware

The i18n middleware automatically detects the user's preferred language from the `Accept-Language` header:

```go
import "yourapp/pkg/middleware"

app := fiber.New()
app.Use(middleware.I18nMiddleware)
```

Test with curl:
```bash
# English
curl -H "Accept-Language: en" http://localhost:8080/api/hello

# Vietnamese
curl -H "Accept-Language: vi" http://localhost:8080/api/hello

# Multiple languages (first supported language is used)
curl -H "Accept-Language: en,vi;q=0.9" http://localhost:8080/api/hello

# Unsupported language (falls back to default)
curl -H "Accept-Language: fr" http://localhost:8080/api/hello
```

### Adding a New Language

1. Create a new directory under `i18n/locales/` for your language code (e.g., `fr` for French)
2. Copy the translation files from an existing language
3. Translate the content
4. The language will be automatically detected and supported

Example:
```bash
# Create French translations
mkdir -p i18n/locales/fr/models
cp i18n/locales/en/auth.json i18n/locales/fr/
cp i18n/locales/en/crud.json i18n/locales/fr/
cp i18n/locales/en/models/user.json i18n/locales/fr/models/
```

### Best Practices

1. **Organize Translations**: Group related translations in separate files (e.g., `auth.json`, `crud.json`)
2. **Use Nested Structures**: Utilize nested JSON structures for better organization
3. **Consistent Keys**: Use consistent key naming across all languages
4. **Default Fallback**: Always provide translations for the default locale
5. **Context-Aware**: Use the context to determine the user's preferred language

## Customization

To customize this template for your project:

1. Define your domain models in `internal/model`
2. Create your service interfaces in `internal/service`
3. Implement your repositories in `internal/repository`
4. Add your API definitions in `pb` directory
5. Update the database schema in `migrations`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

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
