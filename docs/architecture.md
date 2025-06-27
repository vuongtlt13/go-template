# Backend Architecture Documentation

## 🏗️ Architecture Overview

The backend of this Go web application template follows **Clean Architecture** principles and **Domain-Driven Design** patterns, providing a scalable and maintainable foundation for enterprise applications.

## 🎯 Architectural Principles

### 1. Clean Architecture

The application is structured in layers with clear separation of concerns:

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

### 2. Dependency Inversion

- High-level modules don't depend on low-level modules
- Both depend on abstractions
- Abstractions don't depend on details

### 3. Single Responsibility

Each component has a single, well-defined responsibility

## 🏛️ Layer Details

### Presentation Layer

#### HTTP Handlers (`internal/handler/`)

```go
type AuthHandler struct {
    cfg     *config.Config
    service service.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    // 1. Request validation
    // 2. Business logic delegation
    // 3. Response formatting
}
```

**Responsibilities:**

- HTTP request/response handling
- Input validation
- Response formatting
- Error handling

#### Schema Layer (`internal/schema/`)

```go
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=1,max=255"`
}

type LoginResponse struct {
    Token string `json:"token"`
}
```

**Responsibilities:**

- Request/Response data structures
- Validation tags
- API contract definition

#### Router Layer (`internal/routes/`)

```go
type AdminRouter struct {
    app *fiber.App
}

func (r *AdminRouter) Register(app *fiber.App) {
    api := r.app.Group("/api")

    // Register common routes
    common.NewHealthRouter().Register(api)
    common.NewAuthRouter(cfg, authService).Register(api)

    // Register feature routes
    userRouter := admin.NewUserRouter()
    userRouter.Register(api)
}
```

**Responsibilities:**

- Route registration
- Middleware application
- API grouping

### Business Layer

#### Service Layer (`internal/service/`)

```go
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

**Responsibilities:**

- Business logic implementation
- Transaction management
- Service orchestration
- Business rule validation

### Data Layer

#### Repository Layer (`internal/repository/`)

```go
type BaseRepository[T any] interface {
    Create(ctx context.Context, entity *T, db *gorm.DB) error
    FindByID(ctx context.Context, id uint64, preloads []string, db *gorm.DB) (*T, error)
    Update(ctx context.Context, entity *T, db *gorm.DB) error
    Delete(ctx context.Context, entity *T, db *gorm.DB) error
    FindAll(ctx context.Context, offset, limit int, preloads []string, db *gorm.DB) ([]T, error)
}
```

**Responsibilities:**

- Data access abstraction
- CRUD operations
- Query optimization
- Database interaction

#### Model Layer (`internal/model/`)

```go
type User struct {
    ID        uint64         `gorm:"primarykey" json:"id"`
    Email     string         `gorm:"size:255;not null;uniqueIndex" json:"email"`
    Password  string         `gorm:"size:255;not null" json:"-"`
    FullName  string         `gorm:"size:255" json:"full_name"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
    Roles     []Role         `gorm:"many2many:user_roles;" json:"roles,omitempty"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**Responsibilities:**

- Domain entity definition
- Business logic methods
- Database mapping
- Validation rules

## 🔧 Core Infrastructure

### Configuration Management (`pkg/config/`)

```go
type Config struct {
    AppMode  string       `env:"APP_MODE" envDefault:"production"`
    SqlDebug bool         `env:"SQL_DEBUG" envDefault:"false"`
    Cors     string       `env:"CORS" envDefault:"*"`
    Server   ServerConfig `envPrefix:"SERVER_"`
    Database DBConfig     `envPrefix:"DB_"`
    JWT      JWTConfig    `envPrefix:"JWT_"`
    Redis    RedisConfig  `envPrefix:"REDIS_"`
    I18n     I18nConfig   `envPrefix:"I18N_"`
}
```

**Features:**

- Environment variable support
- YAML configuration
- Type-safe configuration
- Default values
- Validation

### Database Layer (`pkg/database/`)

```go
func GetDatabase() *gorm.DB {
    // Singleton pattern with connection pooling
    // PostgreSQL with GORM ORM
    // Auto-migration and soft deletes
}
```

**Features:**

- Connection pooling
- Transaction support
- Migration management
- Query logging
- Performance optimization

### Authentication System (`pkg/auth/`)

```go
type JWTManager struct {
    GenerateToken(userID uint64) (string, error)
    VerifyToken(tokenStr string) (uint64, error)
    ValidateToken(tokenStr string) (*Claims, error)
}
```

**Features:**

- JWT token generation/validation
- bcrypt password hashing
- Token expiration management
- Claims extraction

## 🏢 Multi-Server Architecture

### Server Types

#### 1. Admin Server (`cmd/admin.go`)

- **Purpose**: Administrative interface and APIs
- **Port**: 8000 (configurable)
- **Features**: User management, role management, system administration

#### 2. User Server (`cmd/user.go`)

- **Purpose**: User-facing APIs and services
- **Port**: 8001 (configurable)
- **Features**: User authentication, profile management, user-specific features

#### 3. Common APIs

- **Purpose**: Shared functionality between servers
- **Features**: Health checks, authentication, internationalization

### Server Implementation

```go
type AdminServer struct {
    *server.BaseServer
    logger logger.Logger
}

func (s *AdminServer) Start() error {
    app := s.GetApp()

    // Initialize and register admin routes
    adminRouter := routes.NewAdminRouter()
    adminRouter.Register(app)

    // Swagger documentation
    app.Get("/swagger/*", swagger.HandlerDefault)

    // Start server
    addr := fmt.Sprintf(":%d", s.GetConfig().Server.Port)
    return app.Listen(addr)
}
```

## 🔒 Security Architecture

### Authentication Flow

1. **Login Request** → Handler validates input
2. **Service Layer** → Validates credentials
3. **Repository** → Retrieves user from database
4. **Password Verification** → bcrypt comparison
5. **JWT Generation** → Creates signed token
6. **Response** → Returns token to client

### Authorization System

```go
// Role-based access control
type Role struct {
    ID          uint64
    Code        string
    Name        string
    Description string
    Permissions []Permission
}

type Permission struct {
    ID          uint64
    Code        string
    Name        string
    Service     string
    Method      string
}
```

### Security Middleware

```go
func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // 1. Extract token from Authorization header
        // 2. Validate JWT token
        // 3. Extract claims
        // 4. Store in context
        // 5. Continue to next handler
    }
}
```

## 📊 Data Management

### DataTable System (`pkg/datatable/`)

Advanced data listing with comprehensive features:

```go
type DataTable interface {
    Render(c *fiber.Ctx, extra map[string]interface{}) error
    AddColumn(columnName string, producer ProducerFunc) DataTable
    EditColumn(columnName string, producer ProducerFunc) DataTable
    FilterColumn(columnName string, filter FilterFunc) DataTable
}
```

**Features:**

- Pagination
- Search and filtering
- Column customization
- Export functionality (Excel, CSV, PDF)
- Smart search
- Performance optimization

### Migration System

```sql
-- Versioned database migrations
-- Up migrations
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    -- ...
);

-- Down migrations
DROP TABLE users;
```

## 🌐 Internationalization

### i18n System (`pkg/i18n/`)

```go
func T(ctx context.Context, key string) string {
    // 1. Extract locale from context
    // 2. Load translations from JSON files
    // 3. Return translated text
    // 4. Fallback to default locale
}
```

**Features:**

- JSON-based translation files
- Context-aware locale detection
- Nested key support
- Fallback mechanisms
- Dynamic loading

## ⚡ Performance Features

### Connection Pooling

```go
sqlDB.SetMaxIdleConns(50)
sqlDB.SetMaxOpenConns(200)
sqlDB.SetConnMaxLifetime(30 * time.Minute)
```

### Caching Strategy

- **Redis**: Session storage and data caching
- **In-memory**: Translation caching
- **Query caching**: Database query results

### Background Processing

```go
func ScheduleCleanup(userService service.UserService) *cron.Cron {
    c := cron.New()

    // Run every day at midnight
    c.AddFunc("0 0 * * *", func() {
        userService.CleanupUnverifiedUsers(ctx, days)
    })

    c.Start()
    return c
}
```

## 🛠️ Development Tools

### Code Generator (`pkg/generator/`)

Automated code generation for CRUD operations:

```bash
# Generate all layers for a model
go run pkg/generator/main.go --cmd=generate --model=User --component=repository
go run pkg/generator/main.go --cmd=generate --model=User --component=service
go run pkg/generator/main.go --cmd=generate --model=User --component=handler
```

**Generated Components:**

- Repository layer
- Service layer
- Handler layer
- Schema layer
- DataTable layer

### API Documentation

Auto-generated Swagger documentation:

```go
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body schema.LoginRequest true "Login credentials"
// @Success 200 {object} schema.LoginResponse
// @Router /api/v1/auth/login [post]
```

## 📈 Scalability Considerations

### Horizontal Scaling

- Stateless application design
- Database connection pooling
- Redis for session management
- Load balancer ready

### Vertical Scaling

- Efficient memory usage
- Optimized database queries
- Caching strategies
- Background job processing

### Monitoring & Observability

- Structured logging with Zap
- Health check endpoints
- Performance metrics
- Error tracking

## 🔄 Data Flow

### Request Flow

```
Client Request → Router → Middleware → Handler → Service → Repository → Database
                                                                           ↓
Client Response ← Router ← Middleware ← Handler ← Service ← Repository ← Database
```

### Authentication Flow

```
Login Request → Auth Handler → Auth Service → User Repository → Database
                                                                      ↓
JWT Token ← Auth Handler ← Auth Service ← User Repository ← Database
```

### DataTable Flow

```
DataTable Request → Handler → DataTable Service → Repository → Database
                                                                      ↓
Paginated Response ← Handler ← DataTable Service ← Repository ← Database
```

## 🎯 Best Practices

### Code Organization

- Clear separation of concerns
- Dependency injection
- Interface-based design
- Error handling patterns

### Security

- Input validation
- SQL injection prevention
- XSS protection
- CSRF protection
- Rate limiting

### Performance

- Database query optimization
- Connection pooling
- Caching strategies
- Background processing

### Maintainability

- Consistent naming conventions
- Comprehensive documentation
- Automated testing
- Code generation

This architecture provides a solid foundation for building scalable, maintainable, and secure Go web applications.
