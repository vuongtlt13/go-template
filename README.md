# Go Web Application Template

A modern Go web application template that provides a solid foundation for building scalable web applications with gRPC/Connect.

## Features

- **Core Functionality**
  - JWT-based authentication and authorization
  - User management system
  - Role-based access control
  - Admin dashboard capabilities

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

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Protocol Buffers compiler (protoc)
- Make (for running build commands)
- Docker and Docker Compose

## Project Structure

```
.
├── cmd/                    # Command-line applications
│   ├── admin.go           # Admin server command
│   └── user.go            # User server command
├── internal/              # Private application code
│   ├── domain/           # Business logic and domain models
│   │   ├── handler/      # API handlers
│   │   ├── repository/   # Data access layer
│   │   └── service/      # Business logic layer
│   └── server/           # Server implementations
├── pkg/                   # Public libraries
│   ├── config/           # Configuration management
│   ├── database/         # Database connection
│   └── logger/           # Logging utilities
├── proto/                # Protocol Buffer definitions
├── migrations/           # Database migrations
└── Makefile             # Build and development commands
```

## Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd your-project
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the project root:
   ```env
   # App Configuration
   APP_MODE=development
   SQL_DEBUG=true
   CORS=*

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

4. **Start infrastructure**
   ```bash
   make start-infra
   ```

5. **Run database migrations**
   ```bash
   make migrate-up
   ```

6. **Generate Protocol Buffer code**
   ```bash
   make gen-proto
   ```

## Running the Application

### Admin Server
```bash
go run main.go admin
```

### User Server
```bash
go run main.go user
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

## Customization

To customize this template for your project:

1. Define your domain models in `internal/domain/model`
2. Create your service interfaces in `internal/domain/service`
3. Implement your repositories in `internal/domain/repository`
4. Add your API definitions in `proto` directory
5. Update the database schema in `migrations`

## Environment Variables

The template uses environment variables for configuration. All required variables are listed in the `.env` file template above. The application will automatically load the `.env` file if it exists.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
