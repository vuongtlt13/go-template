# Go Web Application Template
# Project Structure

This project follows a domain-driven design approach with the following structure:
This is a template for a Go web application with user authentication, email verification, and other common features. Built with Fiber, GORM, and other modern Go libraries.

## Features

- User registration and authentication with JWT
- Email verification with background processing
- Role-based access control (user/admin roles)
- Automatic cleanup of unverified users after 7 days
- PostgreSQL database with GORM ORM
- Background task processing with Asynq
- Scheduled tasks with Cron
- Structured logging with Zap
- Clean architecture with dependency injection
- Connect-based API with Protocol Buffers

## Tech Stack

- **Fiber**: Fast and minimalist web framework
- **GORM**: Feature-rich ORM for PostgreSQL
- **JWT**: Authentication using golang-jwt/jwt
- **Viper**: Configuration management
- **Asynq**: Background task processing with Redis
- **Cron**: Scheduled task execution
- **Zap**: Structured logging
- **Connect**: gRPC and REST API generation
- **go-playground/validator**: Input validation
- **Cobra**: Command-line interface

## Prerequisites

- Go 1.23 or later
- PostgreSQL
- Redis (for background tasks)

## Project Structure

```
.
├── api/
│   └── proto/           # Protocol Buffer definitions
├── cmd/
│   └── main.go         # Application entry point
├── config/
│   └── config.yaml     # Configuration file
├── internal/
│   ├── auth/           # Authentication package
│   ├── background/      # Background tasks
│   ├── delivery/       # HTTP handlers
│   ├── models/         # Database models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic
└── pkg/
    ├── config/         # Configuration package
    └── logger/         # Logging package
```

## Configuration

Copy `config/config.yaml.example` to `config/config.yaml` and configure:

```yaml
server:
  port: 8080

db:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "app"
  sslmode: "disable"

jwt:
  secret: "your-secret-key"
  expire_period: "24h"

email:
  host: "smtp.gmail.com"
  port: 587
  username: "your-email@gmail.com"
  password: "your-app-password"
  from: "no-reply@yourdomain.com"

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0

cleanup:
  unverified_user_period: "168h"  # 7 days
```

## Getting Started

### 1. Start Infrastructure with Docker Compose

Start PostgreSQL and Redis using Docker Compose (recommended for local development):

```bash
docker-compose up -d
```

This will start PostgreSQL on port 5432 and Redis on port 6379 with the correct credentials for your config.

### 2. Clone the repository:
```bash
git clone https://github.com/vuongtlt13/template.git
cd template
   ```

2. Install dependencies:
   ```bash
   # Initialize Go modules if not already done
   go mod init yourapp
   
   # Install required packages
   go get -u github.com/spf13/cobra
   go get -u github.com/gofiber/fiber/v2
   go get -u go.uber.org/zap
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/postgres
   go get -u github.com/hibiken/asynq
   go get -u github.com/robfig/cron/v3
   ```

3. Set up the database:
   ```bash
   createdb app
   ```

4. Start Redis:
   ```bash
   redis-server.go
   ```

## Running the Application

The application has two main components that need to be run:

### 1. HTTP Server

Start the HTTP server to handle API requests:

```bash
go run main.go server.go
```

The server will start on `http://localhost:8080` by default.

### 2. Background Job Worker

Start the background job worker to process tasks like cleaning up unverified users:

```bash
go run main.go job
```

### Additional Commands

View available commands:
```bash
go run main.go --help
```

Use a custom config file:
```bash
go run main.go --config path/to/config.yaml server.go
```

## API Endpoints

### Public Endpoints

- `POST /api/v1/auth/login`: User login
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

- `POST /api/v1/auth/register`: User registration
  ```json
  {
    "email": "user@example.com",
    "password": "password123",
    "full_name": "John Doe"
  }
  ```

- `POST /api/v1/auth/verify`: Email verification
  ```json
  {
    "token": "verification-token"
  }
  ```

### Protected Endpoints

Require `Authorization: Bearer <token>` header

- `GET /api/v1/users`: List users (admin only)
- `GET /api/v1/users/:id`: Get user details
- `PUT /api/v1/users/:id`: Update user
  ```json
  {
    "full_name": "Updated Name",
    "role": "admin"  # admin only
  }
  ```
- `DELETE /api/v1/users/:id`: Delete user (admin only)

## Background Tasks

- Email verification: Sends verification emails asynchronously
- User cleanup: Removes unverified users after 7 days

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
