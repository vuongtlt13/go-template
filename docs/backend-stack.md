# Backend Technology Stack

## 🛠️ Technology Overview

This Go web application template uses a carefully selected stack of modern technologies designed for performance, scalability, and developer productivity.

## 📋 Core Technologies

### Web Framework

**Fiber** - Fast HTTP framework for Go

```go
// High-performance HTTP framework
app := fiber.New(fiber.Config{
    AppName:      "Go Template",
    ReadTimeout:  60 * time.Second,
    WriteTimeout: 60 * time.Second,
    IdleTimeout:  120 * time.Second,
})
```

**Key Benefits:**

- ⚡ **High Performance**: Built on top of FastHTTP
- 🚀 **Zero Memory Allocation**: Efficient memory usage
- 🔧 **Express-like API**: Familiar for developers
- 📦 **Rich Middleware**: Built-in middleware ecosystem
- 🛡️ **Security**: Built-in security features

### Database & ORM

**PostgreSQL + GORM**

```go
// Database connection with GORM
dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
    cfg.Database.Password, cfg.Database.Name)

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logMode),
})
```

**Key Benefits:**

- 🗄️ **ACID Compliance**: Full transaction support
- 🔍 **Advanced Queries**: Complex query capabilities
- 📊 **JSON Support**: Native JSON data types
- 🔒 **Security**: Row-level security
- 📈 **Scalability**: Horizontal and vertical scaling

### Authentication & Security

**JWT + bcrypt**

```go
// JWT token generation
claims := Claims{
    RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    },
    UserID: userID,
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString([]byte(secret))
```

**Key Benefits:**

- 🔐 **Stateless**: No server-side session storage
- 🛡️ **Secure**: Industry-standard cryptography
- ⚡ **Fast**: Minimal overhead
- 🔄 **Scalable**: Works across multiple servers

### Configuration Management

**Viper + Environment Variables**

```go
type Config struct {
    AppMode  string       `env:"APP_MODE" envDefault:"production"`
    Server   ServerConfig `envPrefix:"SERVER_"`
    Database DBConfig     `envPrefix:"DB_"`
    JWT      JWTConfig    `envPrefix:"JWT_"`
}
```

**Key Benefits:**

- 🔧 **Flexible**: Multiple config sources
- 🛡️ **Secure**: Environment variable support
- 📝 **Type-safe**: Strong typing
- 🔄 **Hot Reload**: Runtime configuration changes

### Logging

**Uber Zap**

```go
logger := zap.NewProduction()
defer logger.Sync()

logger.Info("Server started",
    zap.String("port", "8000"),
    zap.String("environment", "production"),
)
```

**Key Benefits:**

- ⚡ **High Performance**: Structured logging
- 📊 **Structured**: JSON output format
- 🔍 **Levels**: Multiple log levels
- 📈 **Observability**: Production-ready logging

## 📦 Dependencies Breakdown

### Core Dependencies

| Package                                  | Version  | Purpose            |
| ---------------------------------------- | -------- | ------------------ |
| `github.com/gofiber/fiber/v2`            | v2.52.6  | HTTP framework     |
| `gorm.io/gorm`                           | v1.26.1  | ORM                |
| `gorm.io/driver/postgres`                | v1.5.11  | PostgreSQL driver  |
| `github.com/golang-jwt/jwt/v5`           | v5.2.2   | JWT authentication |
| `github.com/spf13/viper`                 | v1.20.1  | Configuration      |
| `go.uber.org/zap`                        | v1.27.0  | Logging            |
| `github.com/go-playground/validator/v10` | v10.19.0 | Validation         |

### Development Dependencies

| Package                                | Version | Purpose            |
| -------------------------------------- | ------- | ------------------ |
| `github.com/swaggo/swag`               | v1.16.4 | API documentation  |
| `github.com/golang-migrate/migrate/v4` | v4.18.3 | Database migration |
| `github.com/robfig/cron/v3`            | v3.0.1  | Task scheduling    |
| `github.com/stretchr/testify`          | v1.10.0 | Testing            |

## 🏗️ Architecture Components

### 1. HTTP Server Layer

```go
// Base server configuration
type BaseServer struct {
    app *fiber.App
    cfg *config.Config
}

// Middleware stack
app.Use(recover.New())           // Panic recovery
app.Use(logger.New())            // Request logging
app.Use(cors.New(cors.Config{})) // CORS handling
app.Use(limiter.New(limiter.Config{})) // Rate limiting
```

### 2. Database Layer

```go
// Connection pooling
sqlDB.SetMaxIdleConns(50)
sqlDB.SetMaxOpenConns(200)
sqlDB.SetConnMaxLifetime(30 * time.Minute)

// Repository pattern
type BaseRepository[T any] interface {
    Create(ctx context.Context, entity *T, db *gorm.DB) error
    FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*T, error)
    Update(ctx context.Context, entity *T, db *gorm.DB) error
    Delete(ctx context.Context, entity *T, db *gorm.DB) error
}
```

### 3. Business Logic Layer

```go
// Service layer with dependency injection
type AuthService interface {
    Login(ctx context.Context, cred Credential) (string, error)
    Register(ctx context.Context, cred Credential) error
}

type authService struct {
    db         *gorm.DB
    userRepo   repository.UserRepository
    jwtManager authpkg.JWTManagerInterface
}
```

### 4. Presentation Layer

```go
// HTTP handlers with validation
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req schema.LoginRequest
    if err := validator.ValidateRequest(c, &req); err != nil {
        return err
    }

    token, err := h.service.Login(c.Context(), service.Credential{
        Email:    req.Email,
        Password: req.Password,
    })

    return response.SuccessResponse(c, schema.LoginResponse{Token: token}, "ok")
}
```

## 🔧 Development Tools

### Code Generator

```bash
# Generate CRUD code from models
go run pkg/generator/main.go --cmd=generate --model=User --component=repository
go run pkg/generator/main.go --cmd=generate --model=User --component=service
go run pkg/generator/main.go --cmd=generate --model=User --component=handler
```

**Features:**

- 🚀 **Automated CRUD**: Generate all layers from models
- 📝 **Template-based**: Customizable templates
- 🔧 **Type-safe**: Go generics support
- 📊 **Consistent**: Standardized patterns

### Database Migration

```bash
# Run migrations
make migrate-up
make migrate-down

# Create new migration
migrate create -ext sql -dir migrations -seq create_users_table
```

**Features:**

- 📊 **Versioned**: Sequential migration files
- 🔄 **Reversible**: Up and down migrations
- 🛡️ **Safe**: Transaction-based migrations
- 🔍 **Trackable**: Migration history

### API Documentation

```go
// Swagger annotations
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body schema.LoginRequest true "Login credentials"
// @Success 200 {object} schema.LoginResponse
// @Router /api/v1/auth/login [post]
```

**Features:**

- 📚 **Auto-generated**: From code annotations
- 🔍 **Interactive**: Swagger UI
- 📝 **Comprehensive**: Request/response schemas
- 🔄 **Live**: Real-time documentation

## 🛡️ Security Features

### Authentication

- **JWT Tokens**: Stateless authentication
- **bcrypt Hashing**: Secure password storage
- **Token Expiration**: Configurable TTL
- **Refresh Tokens**: Token renewal mechanism

### Authorization

- **Role-Based Access Control (RBAC)**: User roles and permissions
- **Middleware Protection**: Route-level security
- **Permission System**: Granular access control

### Input Validation

```go
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8,max=255"`
}
```

### Security Headers

```go
// CORS configuration
app.Use(cors.New(cors.Config{
    AllowOrigins: allowOrigins,
    AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
}))

// Rate limiting
app.Use(limiter.New(limiter.Config{
    Max:        60,
    Expiration: 1 * time.Minute,
}))
```

## ⚡ Performance Features

### Database Optimization

- **Connection Pooling**: Efficient connection management
- **Query Optimization**: GORM query optimization
- **Indexing**: Strategic database indexing
- **Preloading**: Eager loading of relationships

### Caching Strategy

- **Redis Integration**: Session and data caching
- **In-Memory Caching**: Translation and config caching
- **Query Caching**: Database query result caching

### Background Processing

```go
// Cron job scheduling
c := cron.New()
c.AddFunc("0 0 * * *", func() {
    userService.CleanupUnverifiedUsers(ctx, 30)
})
c.Start()
```

## 🌐 Internationalization

### i18n System

```go
// Multi-language support
func T(ctx context.Context, key string) string {
    locale, ok := ctx.Value("lang").(string)
    if !ok {
        locale = DefaultLocale
    }

    // Load translation from JSON files
    return getTranslation(locale, key)
}
```

**Features:**

- 🌍 **Multi-language**: JSON-based translations
- 🔄 **Context-aware**: Request-based locale detection
- 📝 **Nested Keys**: Hierarchical translation structure
- 🔄 **Fallback**: Default locale fallback

## 📊 Monitoring & Observability

### Logging

- **Structured Logging**: JSON format for production
- **Log Levels**: Debug, Info, Warn, Error
- **Context Logging**: Request context in logs
- **Performance Logging**: Request timing and metrics

### Health Checks

```go
// Health check endpoint
app.Get("/health", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
        "status": "healthy",
        "timestamp": time.Now(),
        "version": "1.0.0",
    })
})
```

### Metrics

- **Request Metrics**: Response times, status codes
- **Database Metrics**: Query performance, connection stats
- **System Metrics**: Memory usage, CPU utilization

## 🔄 Development Workflow

### Local Development

```bash
# Start infrastructure
make start-infra

# Run migrations
make migrate-up

# Start development server
go run main.go admin

# Generate documentation
make gen-swagger
```

### Testing

```go
// Unit tests with testify
func TestAuthService_Login(t *testing.T) {
    // Test implementation
}

// Integration tests
func TestAuthHandler_Login(t *testing.T) {
    // Integration test implementation
}
```

### Code Quality

- **Linting**: golangci-lint configuration
- **Formatting**: gofmt and goimports
- **Testing**: Comprehensive test coverage
- **Documentation**: Auto-generated API docs

## 🚀 Deployment Considerations

### Containerization

```dockerfile
# Multi-stage build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### Environment Configuration

- **Development**: Local development settings
- **Staging**: Pre-production environment
- **Production**: Production-optimized settings

### Scaling

- **Horizontal**: Multiple server instances
- **Vertical**: Resource optimization
- **Database**: Read replicas and sharding
- **Caching**: Distributed caching with Redis

This technology stack provides a robust, scalable, and maintainable foundation for building enterprise-grade Go web applications.
